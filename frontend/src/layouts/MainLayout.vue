<script setup>
    import Navbar from '@/components/Navbar.vue';
    import NavigationDrawer from '@/components/NavigationDrawer.vue';
    import { ref } from 'vue';
    import { RouterView } from 'vue-router';

    const isDrawerOpen = ref(false);
</script>

<template>
    <v-app>
        <NavigationDrawer v-model="isDrawerOpen"/>
        <Navbar @toggle-drawer="isDrawerOpen = !isDrawerOpen"/>
        <v-main>
            <RouterView />
        </v-main>
    </v-app>
</template>
