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
    <div class="flex justify-between">
      <div>
        <label class="input my-7">
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
        <input type="search" class="grow bg-[#FFFEFC]" placeholder="Search" />
        <kbd class="kbd bg-[#FFFEFC] kbd-sm">⌘</kbd>
        <kbd class="kbd bg-[#FFFEFC] kbd-sm">K</kbd>
        </label>
      </div>
      
      <div>
      
      </div>
    </div>
  </div>
   

  </template>
  