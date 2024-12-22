<script setup>

import { ref, defineEmits } from 'vue';
import axios from 'axios';

const inputSearch = ref('')
const emit = defineEmits(['updateDatas']);

const clearInput = (e) => {
    e.preventDefault();
    inputSearch.value = ""
}

const handleSearchMahasiswa = async (e) => {
    e.preventDefault()

    await axios.get(`http://localhost:3000/check?q=${inputSearch.value}`)
    .then((res) => {
        const datas = res.data
        console.log(res)
        // emit('updateDatas', datas);
    })
    .catch((res) => {
        // emit('updateDatas', 1);
        console.error(res);
    });
}

</script>

<style scoped>

.bg-blur {
    background: rgb(231,0,250);
    background: linear-gradient(90deg, rgba(231,0,250,1) 0%, rgba(44,65,245,1) 50%, rgba(255,255,255,1) 100%);
}

</style>

<template>
    <div class="mt-8 relative">
        <form v-on:submit="handleSearchMahasiswa" method="get" class="relative flex items-center w-full">
            <div class="absolute z-[3] left-[20px]">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#b8b8b8" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-globe"><circle cx="12" cy="12" r="10"/><path d="M12 2a14.5 14.5 0 0 0 0 20 14.5 14.5 0 0 0 0-20"/><path d="M2 12h20"/></svg>
            </div>
            <input type="text" v-model="inputSearch" class="w-full focus:outline-none px-16 py-3 rounded-full bg-slate-900 opacity-[60%] text-white z-[2] relative">
            <a v-on:click="clearInput" class="absolute z-[3] right-[20px] cursor-pointer" :class="inputSearch == '' ? 'opacity-0' : 'opacity-100'">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="#b8b8b8" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-x"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
            </a>

        </form>
        <div class="bg-blur w-full h-[100px] blur-[140px] z-[-1] fixed left-[50%] translate-x-[-50%]"></div>

    </div>
</template>