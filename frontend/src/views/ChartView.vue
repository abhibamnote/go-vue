<template>
    <v-container>
        <div id="container"></div>
    </v-container>
</template>

<script setup>
import { onMounted } from "vue";
import Highcharts from "highcharts/highstock";

onMounted(() => {
    const data = [
        [1685453400000, 176.96, 178.99, 176.57, 177.3],
        [1685539800000, 177.33, 179.35, 176.76, 177.25],
        [1685626200000, 177.7, 180.12, 176.93, 180.09],
        [1685712600000, 181.03, 181.78, 179.26, 180.95],
        [1685971800000, 182.63, 184.95, 178.04, 179.58],
        [1686058200000, 179.97, 180.12, 177.43, 179.21],
        [1686144600000, 178.44, 181.21, 177.32, 177.82],
        [1686231000000, 177.9, 180.84, 177.46, 180.57],
        [1686317400000, 181.5, 182.23, 180.63, 180.96],
        [1686576600000, 181.27, 183.89, 180.97, 183.79],
        [1686663000000, 182.8, 184.15, 182.44, 183.31],
    ];

    Highcharts.stockChart("container", {
        title: {
            text: "Candlestick Chart Example",
        },
        series: [
            {
                type: "candlestick",
                name: "AAPL",
                data: data,
                tooltip: {
                    valueDecimals: 2,
                },
            },
        ],
        lang: {
            locale: 'en'
        }
    });
});
</script>

<style>
#container {
    height: 400px;
    min-width: 310px;
}
</style>
