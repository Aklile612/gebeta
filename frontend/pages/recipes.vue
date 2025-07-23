<script setup>
import prepTime from '~/components/prepTime.vue';
import search from '~/components/search.vue';
</script>

<template>
    <div class="bg-white">
        <!-- title -->
        <div class="h-[20vh]">
            <p class="text-xl font-bold">All Recipes</p>
            <p class="text-sm">Browes meals shared by the users around the world</p>
        </div>



        <!-- main recipe part -->
         <div class="bg-[#c7c2c2] w-full h-[80vh]">
            <div>
                <Search/>
            </div>
            <div>
                <prepTime/>
            </div>
         </div>
    </div>

</template>