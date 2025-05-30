<script setup>
import { useAuthStore } from "@/store/authStore";
import { RouterLink } from "vue-router";
const auth = useAuthStore();
defineEmits(["toggle-drawer"]);
const items = [
    { title: "Chart", route: "/chart" },
    { title: "Option Chain", route: "/option-chain" },
];
</script>

<template>
    <v-toolbar dense dark color="primary">
        <!-- Mobile menu button -->
        <v-btn @click="$emit('toggle-drawer')" icon class="hidden-md-and-up">
            <font-awesome-icon :icon="['fas', 'bars']" size="2x" />
        </v-btn>

        <!-- Logo -->
        <v-toolbar-title class="pl-4">
            <RouterLink
                to="/"
                style="cursor: pointer"
                class="router-link-exact-active router-link-active"
            >
                <v-img
                    src="https://opstra.definedge.com/img/opstrabydefinedge.85b3da37.png"
                    contain
                    height="50"
                    width="199"
                ></v-img>
            </RouterLink>
        </v-toolbar-title>

        <!-- Spacer -->
        <v-spacer></v-spacer>

        <!-- Navigation links -->
        <v-toolbar-items class="hidden-sm-and-down">
            <v-btn text to="/" class="text--white v-btn--active">
                <font-awesome-icon :icon="['fas', 'home']" size="2x" />
                <span class="pl-4"> Home </span>
            </v-btn>
            <v-btn
                v-if="!auth.isAuthenticated"
                text
                to="/pricing"
                class="text--white"
            >
                <font-awesome-icon
                    :icon="['fas', 'money-bill-wave']"
                    size="2x"
                />
                <span class="pl-4"> Plans </span>
            </v-btn>
            <v-btn
                v-if="!auth.isAuthenticated"
                text
                href="/blog/"
                target="_blank"
            >
                <font-awesome-icon :icon="['fab', 'blogger']" size="2x" />
                <span class="pl-4"> Blog </span>
            </v-btn>
            <v-btn v-if="auth.isAuthenticated" text to="/blog/">
                <font-awesome-icon :icon="['fas', 'table']" size="2x" />
                <span class="pl-4"> Portfolio </span>
            </v-btn>
            <v-menu v-if="auth.isAuthenticated" open-on-hover>
                <template v-slot:activator="{ props }">
                    <v-btn v-bind="props"> 
                        Futures 
                        <font-awesome-icon :icon="['fas', 'caret-down']" size="2x" class="pl-4"/>
                    </v-btn>
                </template>

                <v-list>
                    <v-list-item
                        v-for="(item, index) in items"
                        :key="index"
                        :value="index"
                    >
                        <RouterLink :to="item.route" style="color: black; text-decoration: none;">{{item.title}}</RouterLink>
                    </v-list-item>
                </v-list>
            </v-menu>
            <v-menu v-if="auth.isAuthenticated" open-on-hover>
                <template v-slot:activator="{ props }">
                    <v-btn v-bind="props"> 
                        Pair Trading 
                        <font-awesome-icon :icon="['fas', 'caret-down']" size="2x" class="pl-4"/>
                    </v-btn>
                </template>

                <v-list>
                    <v-list-item
                        v-for="(item, index) in items"
                        :key="index"
                        :value="index"
                    >
                        <RouterLink :to="item.route" style="color: black; text-decoration: none;">{{item.title}}</RouterLink>
                    </v-list-item>
                </v-list>
            </v-menu>
        </v-toolbar-items>

        <!-- Login/Sign up -->
        <v-toolbar-items class="hidden-sm-and-down">
            <v-btn v-if="!auth.isAuthenticated" to="/auth/login" text>
                <font-awesome-icon :icon="['fas', 'unlock-alt']" size="2x" />
                <span class="pl-4"> Login/Sign up </span>
            </v-btn>
            <v-btn v-if="auth.isAuthenticated" @click="auth.logout()" text>
                <font-awesome-icon :icon="['fas', 'sign-out-alt']" size="2x" />
            </v-btn>
        </v-toolbar-items>
    </v-toolbar>
</template>

<style></style>
