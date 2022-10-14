<template>
    <div class="bg-white h-14 w-full flex items-center p-4">
        <div class="w-full flex justify-between items-center">
            <div class="flex items-center h-14">
                <RouterLink :to="link.to" v-for="link, idx in links"
                    :class="`h-14 w-14 p-4 border-b-2 ${link.active ? 'border-green-400 bg-green-100 hover:bg-green-200' : 'hover:bg-zinc-100 border-white hover:border-zinc-400'}`"
                    :key="idx">
                    <component :is="link.icon" />
                </RouterLink>
            </div>
            <div class="flex items-center">
                <RouterLink to="/auth/login">
                    Login
                </RouterLink>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ComputerDesktopIcon as CDIBold } from '@heroicons/vue/24/solid';
import { AcademicCapIcon as ACI, ChatBubbleLeftIcon as CBLI, DocumentCheckIcon as DCI, UsersIcon as UI } from '@heroicons/vue/24/outline';
import { useRouter } from 'vue-router';
import { computed } from 'vue';

const router = useRouter();

const links = computed(() => [
    {
        to: '/',
        icon: CDIBold,
        active: router.currentRoute.value.path === '/',
    },
    {
        to: '/classes',
        icon: ACI,
        active: router.currentRoute.value.path === '/classes',
    },
    {
        to: '/messaging',
        icon: CBLI,
        active: router.currentRoute.value.path === '/messaging',
    },
    {
        to: '/grades',
        icon: DCI,
        active: router.currentRoute.value.path === '/grades',
    },
    {
        to: '/students',
        icon: UI,
        active: router.currentRoute.value.path === '/students',
    },
]);

</script>
