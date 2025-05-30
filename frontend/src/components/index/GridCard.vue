<script setup>
    defineProps({
        image: String,
        title: String,
        subTitle: String,
        href: String
    })
</script>

<template>
    <v-col sm="6" cols="12">
        <v-card variant="elevated">
            <v-card-item>
                <img
                    :src="image"
                    width="100%"
                    alt=""
                />
                <div class="v-card-title">{{title}}</div>
                <div class="v-card-subtitle">
                    {{subTitle}}
                </div>
            </v-card-item>
            <v-card-actions class="flex justify-center">
                <v-btn :to="href" class="text-center font-weight-bold" color="primary">
                    EXPLORE
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-col>
</template>
