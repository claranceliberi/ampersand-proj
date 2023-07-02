<template>
    <form class="bg-white dark:bg-gray-800 shadow-md rounded px-8 pt-6 pb-8 mb-4 text-gray-700 dark:text-gray-300" @submit="submitForm">
        <div class="mb-4">
            <label class="block text-sm font-bold mb-2" for="agent"> Agent </label>
            <input
                id="agent"
                v-model="state.agent"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-100 leading-tight focus:outline-none focus:shadow-outline"
                type="number"
                placeholder="Agent"
            />
        </div>

        <div class="mb-4">
            <label class="block text-sm font-bold mb-2" for="battery_in"> Battery In </label>
            <input
                id="battery_in"
                v-model="state.battery_in"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-100 leading-tight focus:outline-none focus:shadow-outline"
                type="number"
                placeholder="Battery In"
            />
        </div>

        <div class="mb-4">
            <label class="block text-sm font-bold mb-2" for="battery_out"> Battery Out </label>
            <input
                id="battery_out"
                v-model="state.battery_out"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-100 leading-tight focus:outline-none focus:shadow-outline"
                type="number"
                placeholder="Battery Out"
            />
        </div>

        <div class="mb-4">
            <label class="block text-sm font-bold mb-2" for="driver"> Driver </label>
            <input
                id="driver"
                v-model="state.driver"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-100 leading-tight focus:outline-none focus:shadow-outline"
                type="number"
                placeholder="Driver"
            />
        </div>

        <div class="mb-4">
            <label class="block text-sm font-bold mb-2" for="station"> Station </label>
            <input
                id="station"
                v-model="state.station"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-100 leading-tight focus:outline-none focus:shadow-outline"
                type="number"
                placeholder="Station"
            />
        </div>

        <div class="mb-6">
            <label class="block text-sm font-bold mb-2" for="vehicle"> Vehicle </label>
            <input
                id="vehicle"
                v-model="state.vehicle"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 dark:text-gray-100 leading-tight focus:outline-none focus:shadow-outline"
                type="number"
                placeholder="Vehicle"
            />
        </div>

        <div class="flex items-center justify-between">
            <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">Submit</button>
        </div>
    </form>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { TelemetryData } from "../types/index"
import axios from "axios"

const state = reactive({
    agent: "",
    battery_in: "",
    battery_out: "",
    driver: "",
    station: "",
    vehicle: "",
})

// submit form to the api
const submitForm = async (e: Event) => {
    e.preventDefault()

    const response = await axios.post<TelemetryData>("http://localhost:8059/api/v1/swaps/", state)

    if (response.status === 201) {
        alert("Swap created")
    } else {
        alert("Error creating swap")
    }
    console.log(response)
}
</script>
