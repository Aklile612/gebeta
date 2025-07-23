<script setup>
  definePageMeta({
    layout:false
  })
</script>


<template>
  <div class="flex min-h-screen">
    <!-- image background -->
    <div class="hero w-1/2 min-h-screen bg-no-repeat bg-center bg-cover" style="background-image: url('/image/login.png');">
      <div class="flex flex-col justify-end items-center text-neutral-content w-full h-full pb-10 text-center">
        <div class="max-w-md">
          <p class="mt-7">Discover Amazing Recipes</p>
          <p class="mt-7">Join thousands of food lovers sharing their favorite recipes</p>
        </div>
      </div>
    </div>

  
    <!-- Optional Other Side -->
    <div class="w-1/2 bg-white flex items-center justify-center">
      <!-- top part icon and shit -->
       <div>
        <div class="w-16 h-16 bg-orange-500 rounded-full flex items-center justify-center shadow-md">
                  <IconUtensils class="text-white w-8 h-8" />
        </div>

       </div>
       <!-- welcome and login form -->
        <div></div>


    </div>
  </div>
  
  </template>
  