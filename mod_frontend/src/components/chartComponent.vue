<template>
    <div id="chart-bar" class="app-component-column">
        <div class="card-table">
            <mu-row justify-content="center">
                <mu-select class="mu-text-hide" v-model="typeIds" multiple filerable
                           @change="getStatistics()" help-text="Модальности" tags>
<!--            <template slot="selection" slot-scope="scope">-->
<!--                <mu-chip :selected="scope.selected" color="teal">-->
<!--                {{scope.label}}-->
<!--                </mu-chip>-->
<!--            </template>-->
                    <mu-option v-for="type in allTypes" :key="type" :label="type.name" :value="type.id"></mu-option>
                </mu-select>
            </mu-row>

            <mu-row justify-content="center">
                <div class="process-icon" v-show="process">
                    <mu-circular-progress class="demo-circular-progress" :size="52"></mu-circular-progress>
                </div>

                <div class="div-chart-canvas">
                    <canvas class="chartjs-render-monitor" ref="canvasChart"></canvas>
                </div>
            </mu-row>
        </div>
    </div>
</template>

<script>
    import Chart from 'chart.js'
    import axios from 'axios'

    export default {
        name: "chartComponent",
        data: function () {
            return {
                process: true,
                barChart: null,
                allStats: [],
                allTypes: [],
                typeIds: null,
            }
        },
        methods: {
            getStatistics: function() {
                this.allStats = [];
                let requestData = {
                    "type_ids": this.typeIds
                };
                axios.post('/statistic', requestData)
                .then(response => {
                    if (response.status === 200) {
                        this.allStats = response.data.statistic_languages;
                        // console.log(response.data.statistic_languages[0])
                        this.updateChart();
                    }
                })
                .catch(response => {
                    console.log(response);
                });
            },

            getTypes: async function () {
                axios.get('/types')
                .then(response => {
                    if (response.status === 200) {
                        this.allTypes = response.data.types;
                        this.typeIds = response.data.types.map(obj => obj.id);
                        this.getStatistics();
                    }
                })
                .catch(response => {
                    console.log(response);
                });
            },

            updateChart: function () {
                if (this.barChart) {
                  this.barChart.destroy();
                }
                this.createChart();
            },

            createChart: function () {
                this.process = true;
                this.barChart = new Chart(this.$refs.canvasChart, {
                    type: 'bar',
                    data: {
                        labels: [this.allStats[0].name, this.allStats[1].name],
                        datasets: [{
                            label: 'Частота встречаемости',
                            backgroundColor: [
                                'rgba(252,168,62,0.1)',
                                'rgba(255,99,132,0.1)',
                            ],
                            borderColor: [
                                '#fca83a',
                                'rgb(255,99,132)'
                            ],
                            hoverBackgroundColor: [
                                'rgba(252,168,62,0.4)',
                                'rgba(255,99,132,0.4)',
                            ],
                            pointBackgroundColor: 'white',
                            borderWidth: 1,
                            pointBorderColor: '#249EBF',
                            data: [this.allStats[0].avg_count, this.allStats[1].avg_count]
                        }],
                    },
                    options: {
                        scales: {
                            yAxes: [{
                                ticks: {
                                    beginAtZero: true
                                },
                                gridLines: {
                                    display: true
                                }
                            }],
                            xAxes: [{
                                ticks: {
                                    beginAtZero: true
                                },
                                gridLines: {
                                    display: false
                                }
                            }]
                        },
                        legend: {
                            display: false
                        },
                        responsive: true,
                        maintainAspectRatio: false,
                        height: 200
                    }
                });
                this.process = false;
            }
        },
        mounted() {
            this.getTypes();
        }
    }
</script>

<style scoped>
    .mu-input {
        margin-right: 0px;
    }
</style>