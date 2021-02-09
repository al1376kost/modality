<template>
    <div id="app-modality" class="app-component">
        <div class="left-item">
            <ul class="legend-li" v-if="currentText">
                <li v-for="type in allTypes" :key="type">
                    {{ type.name }}
                </li>
            </ul>
        </div>
        <div class="center-item">
            <mu-row class="bottom-margin" justify-content="between" align-items="center">
                <mu-select v-model="currentLangId" style="width: 100px;">
                    <mu-option v-for="lang in allLanguages" :key="lang"
                               :label="lang.name" :value="lang.id">
                    </mu-option>
                </mu-select>

                <mu-text-field v-if="editMode"
                        class="url-input" v-model="currentUrl" type="text" help-text="URL"></mu-text-field>
                <a class="url-link" v-if="!editMode" :href="this.currentUrl" target="_blank">{{ this.currentUrl }}</a>
            </mu-row>

            <mu-row class="bottom-margin" @contextmenu="openMenu">
                <textarea v-model="currentText" id="txt" class="mod-text"></textarea>
            </mu-row>

            <mu-row class="bottom-margin">
                <div class="left-part-row">
                    <span class="red-text">{{errResult}}</span>
                </div>
                <div class="right-part-row">
                    <button v-if="editMode" class="def-button" @click.prevent="fixText">Сохранить</button>
                </div>
            </mu-row>

            <mu-row justify-content="center" align-items="center">
                <button class="def-button" @click="openTexts=true">Список текстов</button>
            </mu-row>

            <ul id="right-click-menu" tabindex="-1" ref="right" v-if="viewMenu"
                @blur="closeMenu" :style="{top:top, left:left}">
                <li @click="deleteModalityClick">удалить модальность</li>
                <li v-for="type in allTypes" :key="type" @click="chooseType(type); selectHighlightedText()">
                    {{ type.name }}
                </li>
            </ul>
        </div>

<!--        <div class="right-item">-->
<!--            <mu-row>-->
<!--                <button class="text-button" @click="openTexts=true">Список текстов</button>-->
<!--            </mu-row>-->
<!--        </div>-->

        <text-page-component v-if="openTexts" :open="openTexts"
                @getText="getText" @addText='nextText' @close="openTexts=false">
        </text-page-component>
    </div>
</template>

<script>
    /* eslint-disable no-unused-vars */

    import axios from 'axios';
    import $ from 'jquery';

    export default {
        name: "modalityComponent",
        components: {
            textPageComponent: () => import('./textPageComponent.vue')
        },
        data: function () {
            return {
                textData: {
                    text: null,
                    url: null,
                    lang: {
                        id: null
                    }
                },
                selectedText: '',
                selectedModalityStart: null,

                currentText: null,
                currentUrl: null,
                currentTextId: null,
                currentModalities: null,
                modalitiesObjectArray: null,

                viewMenu: false,
                top: '0px',
                left: '0px',

                typeChoice: null,
                allTypes: null,

                allLanguages: null,
                currentLangId: null,

                openTexts: false,
                textHighlighted: false,

                editMode: true,
                // fixMode: false,

                modalitiesColors: ['green', 'light-red', 'blue', 'yellow', 'violet', 'gray', 'mint', 'purple', 'ginger',
                    'peach', 'brown', 'pink', 'light-blue', 'orange', 'red', 'acid-green', 'fluorescent-orange',
                    'prune'],

                errResult: null,
            }
        },
        methods: {
            fixText: function () {
                if (this.currentLangId) {
                    if (!this.currentTextId) {
                        // this.fixMode = true;
                        this.putText();
                    } else if (this.currentTextId) {
                        // this.fixMode = true;
                        this.patchText();
                    }
                } else {
                    this.throwErrorMessage('Выберите язык данного текста');
                }
            },

            nextText: function () {
                this.typeChoice = null;
                this.currentText = null;
                this.currentUrl = null;
                this.currentTextId = null;
                this.currentLangId = null;
                this.destroyTextHighlight(this.textHighlighted);
                document.getElementsByTagName("textarea")[0].readOnly = false;
                // this.fixMode = false;
                this.editMode = true;
            },

            selectHighlightedText: function () {
                let txtArea = document.getElementById("txt");
                this.selectedText = txtArea.value.substring(txtArea.selectionStart, txtArea.selectionEnd);
                this.selectedModalityStart = txtArea.selectionStart;

            },

            chooseType: function (type) {
                if (this.currentTextId && !this.editMode) {
                    this.typeChoice = type;
                    this.putModality();
                } else {
                    this.throwErrorMessage('Сохраните текст');
                }
                this.closeMenu();
            },

            checkModalityApposition: function(startSymbol, length) {
                if(this.modalitiesObjectArray) {
                    for (let obj of this.modalitiesObjectArray) {
                        if (startSymbol >= obj.start_symbol &&
                            startSymbol <= obj.start_symbol + obj.text.length - 1) {
                            this.deleteModality(obj.id);
                        } else if (startSymbol <= obj.start_symbol &&
                            obj.start_symbol <= startSymbol + length) {
                            this.deleteModality(obj.id);
                        }
                    }
                }
            },

            putModality: function () {
                if (this.typeChoice && this.selectedText) {
                    let requestData = {
                        text: this.selectedText,
                        type_id: this.typeChoice.id,
                        text_id: this.currentTextId,
                        start_symbol: this.selectedModalityStart
                    };
                    this.checkModalityApposition(requestData.start_symbol, requestData.text.length - 1);
                    axios.put('/modality', requestData)
                    .then(response => {
                        if (response.status === 200) {
                            this.getText(this.currentTextId, false);
                            this.cancelPutModality();
                        }
                    })
                    .catch(error => {
                        // console.log(error.response.data.error);
                        this.throwErrorMessage('Ошибка добавления');
                    })
                }
            },

            deleteModality: function(id) {
                axios.delete(`/modality?id=${id}`)
                .then(response => {
                    console.log(response);
                })
                .catch(error => {
                    console.log(error.response.data.error);
                })
            },

            deleteModalityClick: function() {
                if (this.selectedModalityStart && this.selectedText) {
                    this.checkModalityApposition(this.selectedModalityStart, this.selectedText.length - 1);
                    this.getText(this.currentTextId, false);
                } else {
                    console.log(this.selectedModalityStart, this.selectedText)
                    this.throwErrorMessage('Выберите модальность');
                }
                this.closeMenu();
            },

            cancelPutModality: function () {
                this.errResult = null;
                this.typeChoice = null;
                this.selectedText = null;
                this.selectedModalityStart = null;
            },

            getLanguages: function () {
            axios.get('/langs')
            .then(response => {
                if (response.status === 200) {
                    this.allLanguages = response.data.languages;
                }
            })
            .catch(error => {
                console.log(error.response.data.error);
            })
            },

            getTypes: function () {
                axios.get('/types')
                .then(response => {
                    this.allTypes = response.data.types;
                })
                .catch(error => {
                    console.log(error.response.data.error);
                })
            },

            getText: function (textId, isEdit) {
                this.editMode = false;
                this.currentLangId = null;
                axios.get(`text?id=${textId}`)
                .then(response => {
                    this.currentText = response.data.text;
                    this.currentUrl = response.data.url;
                    this.currentTextId = response.data.id;
                    this.currentLangId = response.data.lang.id;
                    if (isEdit) {
                        this.editMode = true;
                    }
                })
                .then(() => {
                    this.getCurrentTextModalities(textId);
                })
                .catch(error => {
                    console.log(error.response.data.error);
                })
            },

            getCurrentTextModalities: function (textId) {
                axios.get(`modalities?id=${textId}`)
                .then(response => {
                    this.modalitiesObjectArray = response.data.modalities;
                    // this.currentModalities = response.data.modalities.map(o => o.text);
                })
                .then(() => {
                    this.destroyTextHighlight(this.textHighlighted);
                    let indexArray = this.splitModalitiesByType();
                    if (indexArray) {
                        this.highlightText(indexArray);
                    }
                })
                .catch(error => {
                    console.log(error.response.data.error);
                })
            },

            putText: function () {
                let requestData = {
                    text: this.currentText,
                    url: this.currentUrl,
                    lang: {
                        id: this.currentLangId
                    }
                }
                axios.put('/text', requestData)
                    .then(response => {
                        if (response.status === 200) {
                            this.currentTextId = response.data.id;
                            this.fixEditMode();
                        }
                    })
                    .catch(error => {
                        this.throwErrorMessage(error.response.data.error);
                    });
            },

            patchText: function () {
                let requestData = {
                    id: this.currentTextId,
                    text: this.currentText,
                    url: this.currentUrl,
                    lang: {
                        id: this.currentLangId
                    }
                }
                axios.patch('/text', requestData)
                    .then(response => {
                        this.currentTextId = response.data.id;
                        this.fixEditMode();
                    })
                    .catch(error => {
                        if (error.response.status === 422) {
                            if (error.response.data.error !== 'old and new datas is equal') {
                                this.throwErrorMessage(error.response.data.error);
                            } else {
                                this.editMode = false;
                            }
                        }
                    });
            },

            highlightText: function (highlight) {
                this.textHighlighted = true;
                $('.mod-text').highlightWithinTextarea({
                    highlight: highlight
                })
            },

            destroyTextHighlight: function (isHighlighted) {
                if (isHighlighted) {
                    $('.mod-text').highlightWithinTextarea('destroy');
                    this.textHighlighted = false;
                }
            },

            splitModalitiesByType: function () {
                if (this.modalitiesObjectArray) {
                    let splittedModalitiesArray = [];
                    let tmpIndexArray = [];
                    for (let i = 1; i < 19; i++) {
                        for (let obj of this.modalitiesObjectArray) {
                            if (obj.type_id === i) {
                                let lastSymbol = obj.start_symbol + obj.text.length;
                                tmpIndexArray.push([obj.start_symbol, lastSymbol]);
                            }
                        }
                        splittedModalitiesArray.push(
                            {
                                'highlight': tmpIndexArray,
                                'className': this.modalitiesColors[i - 1]
                            }
                        );
                        tmpIndexArray = [];
                    }
                    return splittedModalitiesArray;
                } else return null
            },

            setMenu: function(top, left) {

                let largestHeight = window.innerHeight - this.$refs.right.offsetHeight - 25;
                let largestWidth = window.innerWidth - this.$refs.right.offsetWidth - 25;

                if (top > largestHeight) top = largestHeight;

                if (left > largestWidth) left = largestWidth;

                this.top = top + 'px';
                this.left = left + 'px';
            },

            closeMenu: function() {
                this.viewMenu = false;
            },

            openMenu: function(e) {
                this.selectHighlightedText();
                this.viewMenu = true;

                this.$nextTick(function() {
                    this.$refs.right.focus();
                    this.setMenu(e.y, e.x);
                }.bind(this));
                e.preventDefault();
            },

            throwErrorMessage: function (message) {
                this.errResult = message;
                const self = this;
                    setTimeout(function () {
                    self.errResult = null;
                }, 2000);
            },

            fixEditMode: function () {
                document.getElementById("txt").readOnly = "true";
                this.editMode = false;
            },
        },
        mounted() {
            this.getLanguages();
            this.getTypes();
        }
    }
</script>

<style scoped>

</style>