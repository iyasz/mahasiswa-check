<script setup>
  import Search from './components/Search.vue'
  import { ref } from 'vue';

  const datas = ref(null);
  const displayedDatas = ref([]); 
  const remainingDatas = ref([]);
  const isLoad = ref(false);
  const isDataNull = ref(false);


  const handleUpdateDatas = (newDatas) => {

    if(newDatas == 1){
      isDataNull.value = true
      return false
    }

    datas.value = newDatas;
    displayedDatas.value = datas.value.slice(0, 50);

    if (datas.value.length > 50) {
      isLoad.value = true;
      remainingDatas.value = datas.value.slice(50);
    } else {
      isLoad.value = false;
      remainingDatas.value = [];
    }
  };


  const handleLoadMore = () => {
    isLoad.value = false;

    const moreDatas = remainingDatas.value.slice(0, 50);
    displayedDatas.value = [...displayedDatas.value, ...moreDatas];

    remainingDatas.value = remainingDatas.value.slice(50);
  }



</script>

<style scoped>
  .fadeup-animation {
    animation: fade-up linear;
    animation-timeline: view();
    animation-range: entry 0% cover 20%;
    overflow: hidden;
  }

  @keyframes fade-up {
    0% {
      opacity: 0;
      transform: translateY(20px);
      scale: 0.9;
    }
    100% {
      opacity: 1;
      transform: translateY(0);
      scale: 1;
    }
  }

</style>

<template>

<div class="flex justify-center ">
  <div class="md:w-[700px] w-[90%]">
    <Search @updateDatas="handleUpdateDatas" />

    <div class="mt-36" v-if="isDataNull == true">
      <h1 class="md:text-[30px] text-3xl font-semibold text-white text-center">Oopss .. Tidak bisa menemukan data</h1>
    </div>

    
    <div v-if="!datas && !isDataNull" class="text-center mt-36">
      <h1 class="md:text-[60px] text-3xl font-semibold text-white flex items-center justify-center md:gap-5 gap-2">Cari Mahasiswa 
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-search md:w-[70px] w-[30px]"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>
      </h1>
      <p class="md:text-xl text-sm text-white md:mt-7 mt-4 opacity-[60%] tracking-widest">Lagi mau nyari nama siapa nih 👀 </p>
    </div>

    <div v-if="datas && !isDataNull" class="my-12" >
      <h1 class="text-3xl font-semibold text-white flex items-center gap-3">
        <svg xmlns="http://www.w3.org/2000/svg" width="27" height="27" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="1.75" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-layout-grid"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/></svg>
        List Mahasiswa</h1>
        <div class="mt-7">

          <div v-for="(data, index) in displayedDatas" :key="index" class="w-full fadeup-animation flex items-center group md:h-[85px] h-full py-4 bg-slate-700/50 z-[3] mb-4 rounded-md cursor-pointer">
            <div class="mx-4 flex items-center w-full gap-6">
              <div class="md:block hidden">
                <div class="flex items-center justify-center w-[30px] h-[30px] border-[3px] border-slate-500 rounded-full">
                  <div class="w-[15px] h-[15px] bg-slate-500 rounded-full hidden group-hover:block"></div>
                </div>
              </div>  
              <div class="">
                  <h2 class="text-white font-medium tracking-wider mb-2 leading-[1.2]">{{ data.nama }} <span class="ms-1 text-xs md:hidden inline"> ({{ data.nim }})</span></h2>
                  <p class="md:text-sm text-xs text-white opacity-75 ">{{ data.nama_pt }} - {{ data.nama_prodi }}</p>
              </div>
              <div class="ms-auto md:block hidden">
                <div class="bg-slate-500 rounded-full px-3 py-1 text-xs text-white opacity-[80%]">
                  {{ data.nim }}
                </div>
              </div>
            </div>
          </div>
          
          <div v-if="isLoad" class="flex justify-center mt-10">
            <button v-on:click="handleLoadMore" class="bg-slate-600 text-white/75 px-5 py-1 text-sm font-semibold rounded-full ">Load More</button>
          </div>
        </div>
    </div>
    
  
  </div>
</div>


</template>

