<template>
    <div class="max-w-sm w-full mx-auto">
        <h1 class="font-bold text-2xl text-center mb-4">Sign in</h1>
        <Card class="p-8">
            <form class="flex flex-col gap-4 mt-4">
                <fieldset class="flex flex-col gap-2">
                    <label for="username" class="text-sm text-slate-600">Username</label>
                    <input
                        class="px-4 py-2 bg-slate-50 border border-slate-300 focus:border-green-300 rounded outline-none focus:ring-4 ring-green-200"
                        id="username" name="username" type="text" placeholder="john.doe" />
                </fieldset>

                <fieldset class="flex flex-col gap-2">
                    <label for="password" class="text-sm text-slate-600">Password</label>
                    <input
                        class="px-4 py-2 bg-slate-50 border border-slate-300 focus:border-green-300 rounded outline-none focus:ring-4 ring-green-200"
                        id="password" name="password" type="password" placeholder="n0tmyp4ssw0rd!" />
                </fieldset>

                <div class="flex justify-center">
                    <button @click="loginClick"
                        class="bg-green-500 hover:bg-green-600 focus:ring-green-400 focus:ring-2 focus:outline-none focus:ring-opacity-50 disabled:opacity-50 transition duration-100 text-white rounded shadow-sm py-2 px-4">
                        Sign in
                    </button>
                </div>
            </form>
        </Card>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import Card from '../../components/Card.vue';

const router = useRouter();

const loginClick = (event: MouseEvent) => {
    router.push('/');
};
</script>