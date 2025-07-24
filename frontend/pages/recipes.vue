<script setup>
import prepTime from '~/components/prepTime.vue';
import search from '~/components/search.vue';
import ratingFilter from '~/components/ratingFilter.vue';
</script>

<template>
    <div class="bg-[#fae3cd] pt-[70px] flex flex-col justify-center">
      
      <!-- title -->
      <div class="h-[10vh] flex flex-col justify-center items-center">
        <p class="text-2xl font-bold">All Recipes</p>
        <p class="text-sm">Browse meals shared by users around the world</p>
      </div>
  
      <!-- main recipe part -->
      <div class="bg-[#ebe9e9] py-6">
        
        <!-- search bar -->
        <div class="flex justify-center mb-4">
          <label class="input bg-white border-black w-full max-w-xl flex items-center gap-2 px-4 py-2 rounded">
            <svg class="h-[1em] opacity-50" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
              <g stroke-linejoin="round" stroke-linecap="round" stroke-width="2.5" fill="none" stroke="currentColor">
                <circle cx="11" cy="11" r="8"></circle>
                <path d="m21 21-4.3-4.3"></path>
              </g>
            </svg>
            <input type="search" class="grow bg-[#FFFEFC] outline-none" placeholder="Search Recipes..." />
            <kbd class="kbd bg-[#FFFEFC] kbd-sm">⌘</kbd>
            <kbd class="kbd bg-[#FFFEFC] kbd-sm">K</kbd>
          </label>
        </div>
  
        <!-- filters -->
        <div class="flex justify-center gap-4">
          <prepTime />
          <ratingFilter />
          <!-- <button class="btn btn-neutral btn-outline"><IconFunnel /> More</button> -->
        </div>
        
      </div>
    </div>
  </template>
