<template>
    <div id="pagination">
        <div class="filter_user_radio radio_div pagination" v-if="pageQuantity>1">
            <span class="pg-num" v-if="currentPage>2"><a @click="previousPage">назад</a></span>

            <label v-for="page in pageQuantity" :key="page" class="pg-num">
<!--                <div v-if="page < currentPage - 1"></div>-->
<!--                <div v-if="page < currentPage + 3"></div>-->
<!--                <div v-else>-->
                <div>
                    <input type="radio" v-model="currentPage" :value="page" @change="update">
                        <span>{{ page }}</span>
                </div>
            </label>

            <span class="pg-num" v-if="currentPage < pageQuantity - 2">
                <a @click="nextPage">вперед</a>
            </span>
        </div>
    </div>
</template>

<script>
    export default {
        name: "pagination",
        props: {
            quantity: {
                type: Number,
                required: true
            },
            limit: {
                type: Number,
                required: true
            },
            currentPage: {
                type: Number,
                required: true
            }
        },
        data: function () {
            return {}
        },
        computed: {
            pageQuantity: function () {
                return Math.ceil(this.quantity / this.limit);
            }
        },
        methods: {
            setPage: function () {
                this.$emit('setPage', this.currentPage);
            },

            nextPage: function () {
                this.currentPage++;
                this.setPage();
                this.$emit('update')
            },

            previousPage: function () {
                this.currentPage--;
                this.setPage();
                this.$emit('update');
            },

            update: function () {
                this.setPage();
                this.$emit('update');
            }
        }
    }
</script>

<style scoped>
    a {
        color: #6f6f6f;
    }
    .pagination {
        margin-top: 15px !important;
        margin-left: 0 !important;
    }
</style>