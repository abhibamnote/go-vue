import AuthLayout from "@/layouts/AuthLayout.vue";
import LoginView from "@/views/LoginView.vue";
import SignupView from "@/views/SignupView.vue";


const routes = [
    {
        path: "/auth",
        component: AuthLayout,
        children: [
            {
                path: "login",
                name: "Login",
                component: LoginView,
            },
            {
                path: "signup",
                name: "Signup",
                component: SignupView,
            },
        ],
    },
];

export default routes;
