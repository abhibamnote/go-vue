<script setup>
import router from "@/router";
import axios from "axios";
import { reactive } from "vue";

const form = reactive({
    firstName: "",
    lastName: "",
    email: "",
    password: "",
    confirmPassword: "",
});

const errors = reactive({
    firstName: false,
    lastName: false,
    email: false,
    password: false,
    confirmPassword: false,
});

const submitForm = async () => {
    // Reset all errors
    Object.keys(errors).forEach((key) => (errors[key] = false));

    // Validate inputs
    let valid = true;
    if (!form.firstName.trim()) {
        errors.firstName = true;
        valid = false;
    }
    if (!form.lastName.trim()) {
        errors.lastName = true;
        valid = false;
    }
    if (!form.email.trim()) {
        errors.email = true;
        valid = false;
    }
    if (!form.password.trim()) {
        errors.password = true;
        valid = false;
    }
    if (form.password !== form.confirmPassword) {
        errors.confirmPassword = true;
        valid = false;
    }

    if (valid) {
        try {
            const response = await axios.post(
                "http://127.0.0.1:8000/api/auth/signup",
                {   
                    firstName: form.firstName,
                    lastName: form.lastName,
                    email: form.email,
                    password: form.password,
                }
            );

            console.log("Login successful:", response.data);
            router.push('/auth/login')
        } catch (err) {
            console.error("Login error:", err);
        }
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
                            <fieldset>
                                <label for="firstName">First Name</label>
                                <input
                                    type="text"
                                    id="firstName"
                                    v-model="form.firstName"
                                    :class="{ 'input-error': errors.firstName }"
                                />
                                <span v-if="errors.firstName" class="error-msg"
                                    >First name is required</span
                                >
                            </fieldset>

                            <fieldset>
                                <label for="lastName">Last Name</label>
                                <input
                                    type="text"
                                    id="lastName"
                                    v-model="form.lastName"
                                    :class="{ 'input-error': errors.lastName }"
                                />
                                <span v-if="errors.lastName" class="error-msg"
                                    >Last name is required</span
                                >
                            </fieldset>

                            <fieldset>
                                <label for="email">Email</label>
                                <input
                                    type="text"
                                    id="email"
                                    v-model="form.email"
                                    :class="{ 'input-error': errors.email }"
                                />
                                <span v-if="errors.email" class="error-msg"
                                    >Email is required</span
                                >
                            </fieldset>

                            <fieldset>
                                <label for="password">Password</label>
                                <input
                                    type="password"
                                    id="password"
                                    v-model="form.password"
                                    :class="{ 'input-error': errors.password }"
                                />
                                <span v-if="errors.password" class="error-msg"
                                    >Password is required</span
                                >
                            </fieldset>

                            <fieldset>
                                <label for="confirmPassword"
                                    >Confirm Password</label
                                >
                                <input
                                    type="password"
                                    id="confirmPassword"
                                    v-model="form.confirmPassword"
                                    :class="{
                                        'input-error': errors.confirmPassword,
                                    }"
                                />
                                <span
                                    v-if="errors.confirmPassword"
                                    class="error-msg"
                                >
                                    Passwords do not match
                                </span>
                            </fieldset>

                            <input
                                type="submit"
                                class="form-btn"
                                value="Sign Up"
                            />
                        </form>

                        <p>
                            Already have an account?
                            <RouterLink to="/auth/login">Login</RouterLink>
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
.error-msg {
    padding-left: 2rem;
}
</style>
