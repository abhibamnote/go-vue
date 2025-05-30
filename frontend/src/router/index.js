import { createRouter, createWebHistory } from "vue-router";
import authRoutes from "./authRoutes";
import mainRoutes from "./mainRoutes";


const routes = [...authRoutes, ...mainRoutes];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

// Auth guard (same as before)
router.beforeEach((to, from, next) => {
    const isAuthenticated = !!localStorage.getItem("token");
    if (to.meta.requiresAuth && !isAuthenticated) {
        next("/auth/login");
    } else {
        next();
    }
});

export default router;
