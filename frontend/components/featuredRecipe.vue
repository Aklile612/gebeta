<script setup>
    defineProps({
        title: String,
        image: String,
        diffculty: String,
        cookTime: Number,
        prepTime:Number,
        rating: Number,
    })
</script>
<!-- #FFF6ED -->

<template>
    <div class="w-full h-[97vh] bg-[#fae3cd]">
        <div class="mx-[30vw] pt-[40px]">
            <h1 class="mb-5 text-4xl ml-[7vw] font-serif  font-bold">Featured Recipes</h1>
            <p class="text-lg font-semibold ml-5 text-gray-500 -mt-2">Discover the most loved recipe from the community</p>
        </div>
        <div class="ml-10 flex gap-10 mt-12">
            <div class="card card-md bg-white w-80 h-100 shadow-sm rounded-xl overflow-hidden">
                <!-- Image -->
                <figure class="">
                  <img
                    class="w-full h-[180px] object-cover rounded-lg"
                    src="/image/hero.png"
                    alt="food image"
                  />
                </figure>
            
                <!-- Body -->
                <div class="card-body p-4 pt-2">
                
                  <!-- Top Row: Rating -->
                  <div class="flex items-center gap-1 text-sm text-yellow-500">
                    <IconStar class="w-4 h-4" />
                    <span class="text-xs font-medium">4.5</span>
                  </div>
              
                  <!-- Title + Badge -->
                  <div class="flex items-center justify-between mt-1">
                    <h2 class="text-md font-bold text-gray-800">Creamy Truffle Pasta</h2>
                    <span class="badge badge-warning text-[10px]">Premium</span>
                  </div>
              
                  <!-- Description -->
                  <p class="text-sm text-gray-600 mt-1">
                    A rich and indulgent pasta dish with truffle cream sauce
                  </p>
              
                  <!-- Footer Row: User + Time -->
                  <div class="flex items-center justify-between mt-1 text-xs text-gray-500">
                    <div class="flex items-center gap-1">
                        <IconUsers class="w-4 h-4 text-gray-500" />
                      <span>By User</span>
                    </div>
                    <div class="flex items-center gap-1">
                      <IconClock9 class="w-4 h-4" />
                      <span>29 min</span>
                    </div>
                  </div>
                </div>
            </div>

            <div class="card card-md bg-white w-96 h-[65vh] rounded-md shadow-sm">
                <figure>
                    <img
                    class="w-[28vw] h-[50vh] object-cover"
                    src="/image/cake1.png"
                    alt="food image" />
                </figure>
                <div class="card-body flex justify-between">
                    <h2 class="card-title">
                        chocklet berry cake
                    </h2>
                    <div class="flex">

                        <div class="text-yellow-400 text-[10px]"  ><IconStar/></div>
                        <p>Decdent berry cake toked with berry cakes
                        </p>
                    </div>
                    <div class="card-actions justify-end">
                    
                    </div>
                </div>
            </div>
            <div class="card card-md bg-white w-96 h-[65vh] shadow-sm">
                <figure>
                    <img
                    class="w-[28vw] h-[50vh] object-cover"
                    src="/image/vegs.png"
                    alt="food image" />
                </figure>
                <div class="card-body ">
                    <h2 class="card-title">
                    Rainbow Buddha Bowl
                    <div class="text-yellow-400 text-[10px]  "><IconStar/></div>
                    </h2>
                    <p>Nutrstios and colorfull bowl packed with fresh ingredients I hope it works</p>
                    <div class="card-actions justify-end">
                        <h1>hello</h1>
                    </div>
                </div>
            </div>
            
        </div>
    </div>
</template>