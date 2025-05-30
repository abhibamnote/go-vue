<script setup>
const items = [
    {
        src: "https://opstra.definedge.com/banner-cta.webp",
        
    },
    {
        src: "https://opstra.definedge.com/banner-cta.webp",
    }
    
];
</script>

<template>
    <section class="py-8">
        <v-container>
            <v-carousel style="height: fit-content;">
                <v-carousel-item style="height: 250px; width: 900px; margin: 0 auto;"
                    v-for="(item, i) in items"
                    :key="i"
                    :src="item.src"
                ></v-carousel-item>
            </v-carousel>    
        </v-container>
    </section>
</template>
