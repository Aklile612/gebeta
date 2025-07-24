<script setup>
import { useForm, useFieldArray } from 'vee-validate'
import * as yup from 'yup'
import { ref, watch } from 'vue'

// Schema
const schema = yup.object({
  title: yup.string().required(),
  description: yup.string().required(),
  prepTime: yup.number().min(1).required(),
  cookTime: yup.number().min(1).required(),
  images: yup.array().min(1, 'At least one image is required'),
  category: yup.string().required(),
  difficulty: yup.string().required(),
  ingredients: yup.array().of(
    yup.object({
      name: yup.string().required(),
      quantity: yup.string().required(),
    })
  ).min(1),
  steps: yup.array().of(
    yup.object({
      description: yup.string().required(),
    })
  ).min(1),
  isPremium: yup.boolean(),
  price: yup.number().when('isPremium', {
    is: true,
    then: (schema) => schema.required().min(1),
    otherwise: (schema) => schema.notRequired()
  })
})

// VeeValidate setup
const { handleSubmit, errors, values, setFieldValue ,meta, touched } = useForm({
  validationSchema: schema,
  initialValues: {
    title: '',
    description: '',
    prepTime: '',
    cookTime: '',
    images: [],
    category: '',
    difficulty:'',
    ingredients: [{ name: '', quantity: '' }],
    steps: [{ description: '' }],
    isPremium: false,
    price: ''
  }
})

const { fields: ingredientFields, push: addIngredient, remove: removeIngredient } = useFieldArray('ingredients')
const { fields: stepFields, push: addStep, remove: removeStep } = useFieldArray('steps')

// File handling
function handleImageUpload(event) {
  const files = Array.from(event.target.files)
  setFieldValue('images', files)
}

// Submission
const onSubmit = handleSubmit((formValues) => {
  console.log('✅ Submitted:', formValues)
})
</script>


<template> 
  <div class="w-full justify-center flex flex-col items-center bg-[#fae3cd] py-[70px]">

    <div class="mx-20 ">
      <div class="flex">
        <NuxtLink to="/" class="text-lg pt-1 font-bold hover:scale-105"><IconChevronLeft /></NuxtLink>
        <h1 class="font-bold text-lg "> Create New Recipe</h1>
      </div>
      <p>Share Your Culiary Creation with the community</p>
    </div>
    <form @submit.prevent="onSubmit" class="space-y-4 shadow-md shadow-slate-600 h-auto   bg-white max-w-3xl px-5 rounded-[7px] mx-auto">
      <div class="font-bold text-lg mt-2">Basic Information</div>
      <!-- Title -->
       <div class="mb-4 w-full flex flex-col">
        <label class="text-xs ml-3 font-semibold">Recipe Title*</label>
        <input v-model="values.title" placeholder="eg. Holiday tibes" class="input w-2xl bg-white placeholder-slate-300 placeholder:font-semibold placeholder:text-[10px]" />
        <span class="text-red-500 text-sm" v-if="touched?.title && errors.title">{{errors.title }}</span>
      </div>
      <!-- Description -->
      <div class="my-2 w-full flex flex-col">
         <label class="text-sm ml-3 font-semibold">Description*</label>
         <textarea v-model="values.description" placeholder="Describe your recipe and what makes it special" class="w-2xl textarea bg-white placeholder:text-[10px] placeholder-slate-300" />
         <span class="text-red-500 text-sm" v-if="touched?.description && errors.description">{{ errors.description }}</span>
      </div>
  
      <!-- Prep & Cook Time -->
      <div class="flex gap-2  flex-col">
        <label class="text-xs ml-8 font-semibold max-w-full">Times needed</label><br>
        <div class="flex gap-1">
        <input v-model="values.prepTime" type="number" placeholder="Prep Time (min)" class="input bg-[#EFEFEF] placeholder-black placeholder:font-semibold placeholder:text-[10px]" />
        <input v-model="values.cookTime" type="number" placeholder="Cook Time (min)" class="input bg-[#EFEFEF] placeholder-black placeholder:font-semibold placeholder:text-[10px]" />
        </div>
      </div>
  
      <!-- Image Upload -->
      <div class="flex flex-col border-1 border-slate-300 border-spacing-1">
      <label class="text-xs font-semibold ml-3 max-w-full">Images</label>

      <!-- Upload box -->
      <div class="relative bg-white rounded-md h-40 w-2xl flex items-center justify-center">
        <div class="text-center space-y-1">
          <p class="text-slate-300 font-semibold">Add your images here</p>
          
          <!-- Button triggers file input -->
          <button
            type="button"
            @click="$refs.imageInput.click()"
            class="px-4 py-2 text-white bg-orange-500 rounded-md font-semibold hover:bg-orange-100 transition"
          >
            Upload Images
          </button>
        </div>

        <!-- Hidden file input -->
        <input
          ref="imageInput"
          type="file"
          multiple
          @change="handleImageUpload"
          class="hidden"
        />
      </div>
      <!-- Error -->
      <span class="text-red-500 text-sm" v-if="touched?.images && errors.images">
        {{ errors.images }}
      </span>
    </div>

      <!-- Category Dropdown -->
      <div>

      <select v-model="values.category" class="input bg-white">
        <option value="">Select Category</option>
        <option value="1">Breakfast</option>
        <option value="2">Lunch</option>
        <option value="3">Dinner</option>
      </select>
      <span class="text-red-500 text-sm" v-if="touched?.category && errors.category" >{{ errors.category }}</span>
      

      <!-- difficulty dropdown -->
       
      <select v-model="values.difficulty" class="input">
        <option value="">Select difficulty</option>
        <option value="1">Easy</option>
        <option value="2">Medium</option>
        <option value="3">Hard</option>
      </select>
      <span class="text-red-500 text-sm" v-if="touched?.difficulty && errors.difficulty" >{{ errors.category }}</span>
    </div>
      

      <!-- Ingredients -->
      <div class="border-2 bg-white border-gray-300 rounded-md p-4 mt-2">
        <label class="text-sm ml-3 text-md font-bold" >Ingredients</label>
        <div v-for="(ingredient, idx) in ingredientFields" :key="ingredient.key" class="flex gap-2 mt-1">
          <input v-model="values.ingredients[idx].name" placeholder="Name" class="input w-2/3" />
          <input v-model="values.ingredients[idx].quantity" placeholder="Quantity" class="input w-1/3" />
          <button type="button" @click="removeIngredient(idx)" class="w-">❌</button>
        </div>
        <button type="button" @click="addIngredient({ name: '', quantity: '' })" class="btn mt-2">+ Add Ingredient</button>
        <span class="text-red-500 text-sm" v-if="touched?.ingredients && errors.ingredients">{{ errors.ingredients }}</span>
      </div>
  
      <!-- Steps -->
      <div>
        <label class="font-semibold">Steps</label>
        <div v-for="(step, idx) in stepFields" :key="step.key" class="flex gap-2 mt-1">
          <textarea v-model="values.steps[idx].description" placeholder="Step Description" class="textarea" />
          <button type="button" @click="removeStep(idx)">❌</button>
        </div>
        <button type="button" @click="addStep({ description: '' })" class="btn mt-2">+ Add Step</button>
        <span class="text-red-500 text-sm" v-if="touched?.steps && errors.steps">{{ errors.steps }}</span>
      </div>
  
      <!-- Premium Toggle -->
      <div class="flex items-center gap-2">
        <label>Premium?</label>
        <input type="checkbox" v-model="values.isPremium" />
      </div>
  
      <!-- Price -->
      <div v-if="values.isPremium">
        <input type="number" v-model="values.price" placeholder="Price (ETB)" class="input" />
        <span class="text-red-500 text-sm" v-if="touched?.price && errors.price">{{ errors.price }}</span>
      </div>
  
    </form>
    <!-- Submit -->
    <button type="submit" class="btn w-full bg-orange-500 text-white">Submit Recipe</button>
  </div>
  </template>
  