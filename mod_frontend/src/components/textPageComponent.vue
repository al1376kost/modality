<template>
    <div id="text-page" class="app-component-column">
        <mu-dialog width="80%" :open.sync="open" :esc-press-close="false" :overlay-close="false">

            <div class="bottom-margin text-filter">
                <mu-row class="bottom-margin" align-items="center" w>
                    <div class="left-part-row">
                        <mu-text-field v-model="searchText" type="text"
                                    help-text="текст" @input="getPageTexts"></mu-text-field>

                        <mu-text-field v-model="searchUrl" type="text"
                                    help-text="url" @input="getPageTexts"></mu-text-field>

                        <a class="create" title="Сброс" @click="resetFilters">
                            <span style="font-size: 17px;"><font-awesome-icon icon="times-circle"/></span>
                        </a>
                    </div>

                    <div class="right-part-row">
                        <button class="add-button" @click.prevent="addText">Добавить текст</button>
                        <a class="close" @click.prevent="close">
                            <span style="font-size: 18px;"><font-awesome-icon icon="times" /></span>
                        </a>
                    </div>
                </mu-row>
            </div>

            <table class="text-table" v-if="listTexts">
                <thead>
                    <tr>
                        <th>Язык</th>
                        <th>Текст</th>
                        <th>URL</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(obj, index) in listTexts" :key="index">
                        <td>{{ obj.lang.name }}</td>
                        <td>{{ obj.text }}</td>
                        <td>
                            <a class="url-link" v-if="obj.url" :href="obj.url" target="_blank">{{ obj.url }}</a>
                            <span v-else>-</span>
                        </td>
                        <td>
                            <a class="create" @click.prevent="passTextId(obj.id)">
                                <font-awesome-icon icon="pencil-alt" title="Редактировать"/>
                            </a>
                            <a class="create" @click.prevent="openDeleteModal(obj.id)">
                                <font-awesome-icon icon="trash-alt" title="Удалить"/>
                            </a>
                        </td>
                    </tr>
                </tbody>
            </table>

            <div v-if="listTexts">
                <pagination v-if="countResult > limit"
                    :quantity="countResult" :limit="limit" :currentPage="page"
                    @update="getPageTexts" @setPage="setPage">
                </pagination>
            </div>

            <mu-row v-if="errSearch" class="error-response" justify-content="center">
                {{errSearch}}
            </mu-row>

            <mu-dialog v-if="textDeleteId" title="Удалить выбранный текст?" width="500" max-width="80%"
                :esc-press-close="false" :overlay-close="false" :open.sync="openDelete">
                <div v-if="resultDelete">
                    <span class="orange-text">{{resultDelete}}</span>
                    <mu-flex class="flex-wrapper" justify-content="end">
                        <button class="del-button" @click.prevent="closeDeleteModal">Закрыть</button>
                    </mu-flex>
                </div>
                <div v-else>
                    <mu-flex class="flex-wrapper" justify-content="end">
                        <button class="del-button" @click.prevent="closeDeleteModal">НЕТ</button>
                        <button class="del-button yes" @click.prevent="deleteText">ДА</button>
                    </mu-flex>
                </div>
            </mu-dialog>
        </mu-dialog>
    </div>
</template>

<script>
    /* eslint-disable no-unused-vars */

    import axios from 'axios';

    export default {
        name: "textPageComponent",
        components: {
            pagination: () => import('./pagination.vue'),
        },
        props: {
            open: {
                type: Boolean,
                required: true
            },
        },
        data: function () {
            return {
                listTexts: null,
                countResult: null,
                page: 1,
                limit: 5,
                searchText: '',
                searchUrl: '',
                errSearch: '',

                textDeleteId: null,
                openDelete: false,
                resultDelete: null
            }
        },
        methods: {
            getPageTexts: function () {
                this.errSearch = '';
                let requestData = {
                    "page": this.page,
                    "limit": this.limit,
                    "sort_by": [
                        {
                            "name":"it.id",
                            "ascending": true
                        },
                        {
                            "name":"it.object_text",
                            "ascending": true
                        }
                    ],
                    "filter": {
                        "text_like": this.searchText,
                        "url_like": this.searchUrl
                    }
                };
                axios.post('/texts', requestData)
                    .then(response => {
                        if (response.data.obect_texts) {
                            this.listTexts = response.data.obect_texts;
                            this.countResult = response.data.count;
                        } else {
                            this.listTexts = null;
                            this.countResult = null;
                            this.errSearch = 'Ничего не найдено';
                        }
                    })
                    .catch(error => {
                        console.log(error.response.data.error);
                    });
            },

            addText: function () {
                this.$emit('addText');
                this.close();
            },

            passTextId: function(id) {
                this.$emit('getText', id, false);
                this.close();
            },

            close: function() {
                this.$emit('close');
            },

            openDeleteModal: function (id) {
                this.textDeleteId = id;
                this.openDelete = true;
            },

            closeDeleteModal: function () {
                this.openDelete = false;
                this.setInitPage();
                this.getPageTexts();
                this.textDeleteId = null;
                this.resultDelete = null;
            },

            deleteText: function () {
                if (this.textDeleteId) {
                    axios.delete(`/text?id=${this.textDeleteId}`)
                    .then(response => {
                        this.resultDelete = 'Выполнено успешно';
                        this.getPageTexts();
                    })
                } else {
                    this.resultDelete = 'Ошибка при удалении :('
                }
            },

            resetFilters: function () {
                this.searchText = '';
                this.searchUrl = '';
                this.getPageTexts();
            },

            setInitPage: function () {
                this.page = 1;
            },

            setPage: function (page) {
                this.page = page;
            }
        },
        mounted() {
            this.getPageTexts();
            this.setInitPage();
        }
    }
</script>

<style scoped>

</style>