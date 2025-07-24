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
const { handleSubmit, errors, values, setFieldValue } = useForm({
  validationSchema: schema,
  initialValues: {
    title: '',
    description: '',
    prepTime: '',
    cookTime: '',
    images: [],
    category: '',
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
  <div class="w-full bg-[#fae3cd]">

    <div class="mx-auto ">
      <div class="flex gap-5">
        <div class="text-lg font-bold"><IconChevronLeft /></div>
        <h1 class="font-bold text-lg "> Create New Recipe</h1>
      </div>
      <p>Share Your Culiary Creation with the community</p>
    </div>
    <form @submit.prevent="onSubmit" class="space-y-4 py-[70px] bg-white max-w-2xl rounded-[5px] mx-auto">
  
      <!-- Title -->
      <input v-model="values.title" placeholder="Title" class="input" />
      <span class="text-red-500 text-sm">{{values.title ? '': errors.title }}</span>
  
      <!-- Description -->
      <textarea v-model="values.description" placeholder="Description" class="textarea" />
      <span class="text-red-500 text-sm">{{ errors.description }}</span>
  
      <!-- Prep & Cook Time -->
      <div class="flex gap-2">
        <input v-model="values.prepTime" type="number" placeholder="Prep Time (min)" class="input" />
        <input v-model="values.cookTime" type="number" placeholder="Cook Time (min)" class="input" />
      </div>
  
      <!-- Image Upload -->
      <input type="file" multiple @change="handleImageUpload" class="input" />
      <span class="text-red-500 text-sm">{{ errors.images }}</span>
  
      <!-- Category Dropdown -->
      <select v-model="values.category" class="input">
        <option value="">Select Category</option>
        <option value="1">Breakfast</option>
        <option value="2">Lunch</option>
        <option value="3">Dinner</option>
      </select>
      <span class="text-red-500 text-sm">{{ errors.category }}</span>
  
      <!-- Ingredients -->
      <div>
        <label class="font-semibold">Ingredients</label>
        <div v-for="(ingredient, idx) in ingredientFields" :key="ingredient.key" class="flex gap-2 mt-1">
          <input v-model="values.ingredients[idx].name" placeholder="Name" class="input" />
          <input v-model="values.ingredients[idx].quantity" placeholder="Quantity" class="input" />
          <button type="button" @click="removeIngredient(idx)">❌</button>
        </div>
        <button type="button" @click="addIngredient({ name: '', quantity: '' })" class="btn mt-2">+ Add Ingredient</button>
        <span class="text-red-500 text-sm">{{ errors.ingredients }}</span>
      </div>
  
      <!-- Steps -->
      <div>
        <label class="font-semibold">Steps</label>
        <div v-for="(step, idx) in stepFields" :key="step.key" class="flex gap-2 mt-1">
          <textarea v-model="values.steps[idx].description" placeholder="Step Description" class="textarea" />
          <button type="button" @click="removeStep(idx)">❌</button>
        </div>
        <button type="button" @click="addStep({ description: '' })" class="btn mt-2">+ Add Step</button>
        <span class="text-red-500 text-sm">{{ errors.steps }}</span>
      </div>
  
      <!-- Premium Toggle -->
      <div class="flex items-center gap-2">
        <label>Premium?</label>
        <input type="checkbox" v-model="values.isPremium" />
      </div>
  
      <!-- Price -->
      <div v-if="values.isPremium">
        <input type="number" v-model="values.price" placeholder="Price (ETB)" class="input" />
        <span class="text-red-500 text-sm">{{ errors.price }}</span>
      </div>
  
      <!-- Submit -->
      <button type="submit" class="btn w-full bg-orange-500 text-white">Submit Recipe</button>
    </form>
  </div>
  </template>
  