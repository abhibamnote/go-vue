<template>
    <v-container>
        <div id="container"></div>
    </v-container>
</template>

<script setup>
import { onMounted } from "vue";
import Highcharts from "highcharts/highstock";
import { useAuthStore } from "@/store/authStore";
import axios from "axios";

const auth = useAuthStore();
onMounted(async () => {

    const response = await axios.get(import.meta.env.VITE_BACKEND_URL+"/api/chart-data", {
        headers: {
            Authorization: auth.token
        }
    });

    const dataTwo = []

    response.data.data.chartData.forEach(ele => {
        const dataEle = [];
        dataEle.push(ele.RateDate)
        dataEle.push(ele.Open);
        dataEle.push(ele.High)
        dataEle.push(ele.Low)
        dataEle.push(ele.Close)
        dataTwo.push(dataEle)
    })

    Highcharts.stockChart("container", {
        title: {
            text: "Candlestick Chart Example",
        },
        series: [
            {
                type: "candlestick",
                name: "AAPL",
                data: dataTwo,
                tooltip: {
                    valueDecimals: 2,
                },
            },
        ],
        plotOptions: {
            candlestick: {
                color: 'red',        
                lineColor: 'darkred', 
                upColor: 'green',
                upLineColor: 'darkgreen'
            },
        },
        lang: {
            locale: 'en'
        },
        yAxis: {
            minPadding: 0.5,
            opposite: false
        },
        xAxis: {
            minPadding: 0.5
        }
    });
});
</script>

<style>
#container {
    height: 600px;
    min-width: 310px;
}
</style>
