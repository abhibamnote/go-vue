<script setup>
import { RouterLink } from 'vue-router';

    defineProps(['modalValue'])
    defineEmits(['update:modalValue'])
    
</script>

<template>
    <v-navigation-drawer v-model="drawer" temporary :width="200">
        <v-list>
            <v-list-item>
                <RouterLink to="/" class="flex ">
                    <font-awesome-icon :icon="['fas', 'home']" size="2x" class="px-4"/>
                    Home
                </RouterLink>
            </v-list-item>
            <v-list-item>
                <v-btn style="width: 100%;" text to="/" class="text--white">
                    <font-awesome-icon :icon="['fas', 'home']" size="2x" />
                    <span class="pl-4"> Home </span>
                </v-btn>
            </v-list-item>
            <v-list-item>
                <v-btn style="width: 100%;" text to="/" class="text--white">
                    <font-awesome-icon :icon="['fas', 'home']" size="2x" />
                    <span class="pl-4"> Home </span>
                </v-btn>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>
