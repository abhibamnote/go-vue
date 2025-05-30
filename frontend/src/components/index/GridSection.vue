<script setup>
import GridCard from "./GridCard.vue";

const data = [
    {
        image: "https://opstra.definedge.com/img/options-algorithm.a497511f.png",
        title: "Options Algorithm",
        subTitle:
            "Quickly find option trading opportunities in the underlying of your interest",
        href: "/options",
    },
    {
        image: "https://opstra.definedge.com/img/options-algorithm.a497511f.png",
        title: "Options Algorithm 2",
        subTitle:
            "Quickly find option trading opportunities in the underlying of your interest",
        href: "/options",
    },
    {
        image: "https://opstra.definedge.com/img/options-algorithm.a497511f.png",
        title: "Options Algorithm 3",
        subTitle:
            "Quickly find option trading opportunities in the underlying of your interest",
        href: "/options",
    },
    {
        image: "https://opstra.definedge.com/img/options-algorithm.a497511f.png",
        title: "Options Algorithm 4",
        subTitle:
            "Quickly find option trading opportunities in the underlying of your interest",
        href: "/options",
    },
];
</script>

<template>
    <section>
        <v-container>
            <div class="text-center">
                <h2>Your first step towards Options Analysis</h2>
                <p>Start analysing and building your options strategies</p>
            </div>

            <v-row>
                <GridCard
                    v-for="ele in data"
                    :title="ele.title"
                    :sub-title="ele.subTitle"
                    :image="ele.image"
                    :href="ele.href"
                />
            </v-row>
        </v-container>
    </section>
</template>
