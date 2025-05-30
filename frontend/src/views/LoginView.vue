<script setup>
import router from "@/router";
import { useAuthStore } from "@/store/authStore";
import axios from "axios";
import { reactive, ref } from "vue";

const auth = useAuthStore();

const form = reactive({
    email: "",
    password: "",
});

const errors = reactive({
    email: false,
    password: false,
});

const submitForm = async () => {
    errors.email = !form.email.trim();
    errors.password = !form.password.trim();
    
    if (!errors.email && !errors.password) {
        // Form is valid, proceed with login
        const response = await axios.post('http://127.0.0.1:8000/api/auth/login', {
            email: form.email,
            password: form.password
        })
        auth.login(response.data.data.token);
        router.push('/')
        console.log(response);
        
    }
};
</script>

<template>
    <section class="fill-height" style="background-color: black">
        <div class="top">
            <img
                src="https://sso.definedge.com/auth/resources/7.0.0/login/definedge/img/logo_def.png"
                alt=""
            />
        </div>

        <v-container>
            <v-row justify="center">
                <v-col cols="12" sm="5" align-self="center">
                    <div class="wrapper">
                        <h1 class="text-center font-weight-thin py-4">
                            Log In To Defineedge Account
                        </h1>

                        <form @submit.prevent="submitForm">
                            <span v-if="errors.email" class="error-msg"
                                >Email is required</span
                            >
                            <span v-if="errors.password" class="error-msg"
                                >Password is required</span
                            >
                            <fieldset>
                                <label for="email">Email</label>
                                <input
                                    type="text"
                                    v-model="form.email"
                                    :class="{ 'input-error': errors.email }"
                                />
                            </fieldset>
                            <fieldset>
                                <label for="password">Password</label>
                                <input
                                    type="password"
                                    v-model="form.password"
                                    :class="{ 'input-error': errors.password }"
                                />
                            </fieldset>

                            <input
                                type="submit"
                                class="form-btn"
                                value="Log In"
                            />
                        </form>

                        <p>
                            Don't have an account?
                            <RouterLink to="/auth/signup">Register</RouterLink>
                        </p>
                    </div>
                </v-col>
            </v-row>
        </v-container>
    </section>
</template>

<style scoped>
.top {
    display: flex;
    justify-content: center;
    padding-top: 1rem;
}
.top img {
    height: 100px;
}
.wrapper {
    background-color: #dddddd;
    margin-top: 2rem;
}
form fieldset {
    display: flex;
    flex-direction: column;
    border: none;
    padding: 0.25rem 2rem;
}
form fieldset label {
    font-size: 14px;
}
form fieldset input {
    display: block;
    width: 100%;
    height: 36px;
    padding: 2px 6px;
    font-size: 12px;
    line-height: 1.66666667;
    color: #363636;
    background-color: #fff;
    background-image: none;
    border: 1px solid #bbb;
    border-radius: 1px;
    transition: border-color ease-in-out 0.15s, box-shadow ease-in-out 0.15s;
}
.form-btn {
    background-color: #0088ce;
    background-image: linear-gradient(to bottom, #39a5dc 0%, #0088ce 100%);
    background-repeat: repeat-x;
    border-color: #00659c;
    color: #fff;
    box-shadow: 0 2px 3px rgba(3, 3, 3, 0.1);
    margin: 0.25rem 2rem;
    padding: 0.25rem 1rem;
    width: 100px;
    text-align: center;
    border-radius: 1rem;
}
.input-error {
    border: 1px solid red;
}
.wrapper p {
    padding: 1rem 2rem;
    font-size: 12px;
}
.error-msg{
    padding-left: 2rem;
}
</style>
