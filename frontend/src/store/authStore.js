// stores/auth.js or stores/auth.ts
import { defineStore } from "pinia";
import { ref } from "vue";

export const useAuthStore = defineStore("auth", () => {
    const token = ref(localStorage.getItem("token") || null);
    const isAuthenticated = ref(!!token.value);

    function login(newToken) {
        token.value = newToken;
        isAuthenticated.value = true;
        localStorage.setItem("token", newToken);
    }

    function logout() {
        token.value = null;
        isAuthenticated.value = false;
        localStorage.removeItem("token");
    }

    return { token, isAuthenticated, login, logout };
});
