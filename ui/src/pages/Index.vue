<template>
    <Card class="px-0 py-0 pt-4">
        <div class="px-8 py-4 mb-4 flex items-center justify-between">
            <div class="flex flex-col">
                <h2 class="text-2xl font-bold">Classes</h2>
                <h4 class="text-sm text-zinc-500">Your child is enrolled in {{classes.length}} class(es).</h4>
            </div>
            <button>Hello</button>
        </div>

        <table class="w-full">
            <thead class="h-12">
                <tr>
                    <th class="pl-8 pr-4 bg-zinc-50 text-zinc-800 font-bold text-left">
                        Name
                    </th>
                    <th class="pl-4 pr-8 bg-zinc-50 text-zinc-800 font-bold text-left">
                        Description
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr :class="['h-12', 'border-zinc-50', idx == (classes.length - 1) ? '' : 'border-b']"
                    v-for=" c, idx in classes" :key="c.id" :id="`class-${c.id}`">
                    <td class="pl-8 pr-4 text-green-600">
                        <RouterLink :to="`/class/${c.id}`">{{c.display_name}}</RouterLink>
                    </td>
                    <td class="pl-4 pr-8">{{c.description}}</td>
                </tr>
            </tbody>
        </table>
    </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useIpcRenderer } from '@vueuse/electron';
import Card from '../components/Card.vue';

const ipcRenderer = useIpcRenderer();

const result = ipcRenderer.invoke<OSPResponse | null>('classes.list');
const classes = computed<{ id: string; name: string; display_name: string; description: string; created_at: string; updated_at: string }[]>(() => {
    if (result.value) {
        return JSON.parse(result.value.body);
    }
    return [];
});
</script>
