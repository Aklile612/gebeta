<script setup>
import prepTime from '~/components/prepTime.vue';
import search from '~/components/search.vue';
import ratingFilter from '~/components/ratingFilter.vue';
</script>

<template>
    <div class="bg-white pt-[70px] ">
        <!-- title -->
        <div class="h-[20vh] flex justify-center">
            <p class="text-xl font-bold">All Recipes</p>
            <p class="text-sm">Browes meals shared by the users around the world</p>
        </div>



        <!-- main recipe part -->
         <div class="bg-[#ebe9e9] w-full h-[80vh]">
            <div class="w-full ml-16">
                <div class="my-10">
                    <label class="input  bg-white border-black">
                    <svg class="h-[1em] opacity-50" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                        <g stroke-linejoin="round"
                            stroke-linecap="round"
                            stroke-width="2.5"
                            fill="none"
                            stroke="currentColor"
                        >
                            <circle cx="11" cy="11" r="8"></circle>
                            <path d="m21 21-4.3-4.3"></path>
                        </g>
                    </svg>
                    <input type="search" class="grow bg-[#FFFEFC]" placeholder="Search Recipes..." />
                    <kbd class="kbd bg-[#FFFEFC] kbd-sm">⌘</kbd>
                    <kbd class="kbd bg-[#FFFEFC] kbd-sm">K</kbd>
                    </label>
                </div>
            </div>
            <div class="w-full flex gap-3">
                <prepTime/>
                <ratingFilter/>
            </div>
         </div>
    </div>

</template>