<script setup>
import { Search } from '#components';
</script>


<template>
  <div>

    <!-- // hero section -->
    <div  class="hero min-h-screen" style="background-image: url('/image/hero2.png');">
        <div class="hero-overlay"></div>
        <div class="hero-content text-neutral-content text-center">
          <div class="max-w-[90vw]">
            <h1 class="mb-5 text-5xl font-bold">Find Your Next Favorite Recipe</h1>
            <p class="mb-5">
              Browse,cook and share amazing meals from everywhere
            </p>
            <button class="btn btn-primary bg-orange-500 border-orange-400">Get Started</button>
          </div>
        </div>
    </div>

    <!-- search and catagories -->
    <div><search/></div>
  </div>
   

  </template>
  