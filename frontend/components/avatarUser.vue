<script setup>

</script>


<template>
    <div class="bg-white  w-full h-[60vh]">
        <div class="mx-[30vw] pt-[40px]">
            <h1 class="mb-5 text-4xl ml-[7vw] font-serif  font-bold">Top Recipe Creators</h1>

        </div>
        <div class="flex gap-5">
            <div class="w-[240px] p-4 bg-[#fae3cd] rounded-xl flex flex-col items-center shadow-sm">
                <!-- Icon in circle -->
                <div class="w-16 h-16 bg-white rounded-full flex items-center justify-center shadow-md">
                  <IconUsers class="text-orange-500 w-8 h-8" />
                </div>
            
                <!-- User info -->
                <div class="mt-4 text-center">
                  <p class="font-bold text-sm text-gray-800">Alana Germy</p>
                  <p class="text-xs text-gray-600 mt-1">35 Recipes Submitted</p>
                
                  <!-- View Button -->
                  <button
                    class="mt-3 text-sm bg-orange-200 hover:bg-orange-500 text-gray-800 hover:text-white font-medium py-1 px-3 rounded-md transition-all duration-200"
                  >
                    View Recipes
                  </button>
                </div>
            </div>
            <div class="w-[300px] h-[200px] bg-[#fae3cd] rounded-[5px]">
                <div class="avatar">
                <div class="ring-primary ring-offset-base-100 w-30 rounded-full ring-2 ring-offset-2">
                    <img src="https://img.daisyui.com/images/profile/demo/spiderperson@192.webp" />
                </div>
                </div>
                <div class="mt-2 ">
                    <p class="font-bold text-sm">Alana Germy</p>
                    <p> 35 Recipes Submiited</p>
                    <button class="btn bg-orange-300 rounded-md btn-soft  hover:bg-orange-500">View Recipes</button>
                </div>
            </div>

            <div class="w-[300px] h-[200px] bg-[#fae3cd] rounded-[5px]">
                <div class="avatar">
                <div class="ring-primary ring-offset-base-100 w-30 rounded-full ring-2 ring-offset-2">
                    <img src="https://img.daisyui.com/images/profile/demo/spiderperson@192.webp" />
                </div>
                </div>
                <div class="mt-2 ">
                    <p class="font-bold text-sm">Alana Germy</p>
                    <p> 35 Recipes Submiited</p>
                    <button class="btn bg-orange-300 rounded-md btn-soft  hover:bg-orange-500">View Recipes</button>
                </div>
            </div>
            <div class="w-[300px] h-[200px] bg-[#fae3cd] rounded-[5px]">
                <div class="avatar">
                <div class="ring-primary ring-offset-base-100 w-30 rounded-full ring-2 ring-offset-2">
                    <img src="https://img.daisyui.com/images/profile/demo/spiderperson@192.webp" />
                </div>
                </div>
                <div class="mt-2 ">
                    <p class="font-bold text-sm">Alana Germy</p>
                    <p> 35 Recipes Submiited</p>
                    <button class="btn bg-orange-300 rounded-md btn-soft  hover:bg-orange-500">View Recipes</button>
                </div>
            </div>
        </div>
    </div>
</template>