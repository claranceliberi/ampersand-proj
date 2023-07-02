import { TelemetryData } from '../types/index';
<script lang="ts" setup>
import { TelemetryData } from "~/types/index"
import { formatTime } from "~/utils/index"

const viewing = ref(0)

const { data: swaps, refresh } = await useFetch<TelemetryData[]>("http://localhost:8059/api/v1/swaps/")

onMounted(() => {
    // if page is rendered without data, fetch them from the server
    if (!swaps.value || swaps.value.length > 0) {
        console.log("REFRESHING SWAPS")
        refresh()
    }
})
</script>

<template>
    <div>
        <div>
            <div class="relative overflow-x-auto">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" class="px-6 py-3">ID</th>
                            <th scope="col" class="px-6 py-3">Agent</th>
                            <th scope="col" class="px-6 py-3">Vehicle</th>
                            <th scope="col" class="px-6 py-3">Station</th>
                            <th scope="col" class="px-6 py-3">Time</th>
                            <th scope="col" class="px-6 py-3">View</th>
                        </tr>
                    </thead>
                    <tbody>
                        <template v-for="swap in swaps" :key="swap.ID">
                            <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                                <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ swap.ID }}</th>
                                <td class="px-6 py-4">{{ swap.Agent.FirstName }} {{ swap.Agent.LastName }}</td>
                                <td class="px-6 py-4">{{ swap.Vehicle.ModelName }} - {{ swap.Vehicle.Identifier }}</td>
                                <td class="px-6 py-4">{{ swap.Station.Location }}</td>
                                <td class="px-6 py-4">{{ formatTime(swap.CreatedAt) }}</td>
                                <td class="px-6 py-4">
                                    <button
                                        type="button"
                                        class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm p-2.5 text-center inline-flex items-center mr-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                                        @click="viewing === swap.ID ? (viewing = 0) : (viewing = swap.ID)"
                                    >
                                        <svg aria-hidden="true" class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                                            <path
                                                fill-rule="evenodd"
                                                d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z"
                                                clip-rule="evenodd"
                                            ></path>
                                        </svg>
                                        <span class="sr-only">Icon description</span>
                                    </button>
                                </td>
                            </tr>
                            <!-- should be displayed only when selected -->
                            <tr v-if="viewing === swap.ID">
                                <td colspan="6">
                                    <!-- swap details -->
                                    <div class="p-3 w-full inline-block bg-white dark:bg-gray-800 text-gray-700 dark:text-gray-200">
                                        <table class="table-fixed">
                                            <tbody>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2 font-bold text-xl">Swap ID:</td>
                                                    <td class="p-2 w-1/2">{{ swap.ID }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Agent Name:</td>
                                                    <td class="p-2 w-1/2">{{ swap.Agent.FirstName }} {{ swap.Agent.LastName }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Vehicle Name:</td>
                                                    <td class="p-2 w-1/2">{{ swap.Vehicle.ModelName }} - {{ swap.Vehicle.Identifier }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Vehicle Identifier:</td>
                                                    <td class="p-2 w-1/2">{{ swap.Vehicle.Identifier }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Station Name:</td>
                                                    <td class="p-2 w-1/2">{{ swap.Station.Location }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Time:</td>
                                                    <td class="p-2 w-1/2">{{ formatTime(swap.CreatedAt) }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">BatteryIn Identifier:</td>
                                                    <td class="p-2 w-1/2">{{ swap.BatteryIn.Identifier }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">BatteryIn SOC:</td>
                                                    <td class="p-2 w-1/2">{{ swap.BatteryInSoc.toFixed(2) }} %</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">BatteryOut Identifier:</td>
                                                    <td class="p-2 w-1/2">{{ swap.BatteryOut.Identifier }}</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">BatteryOut SOC:</td>
                                                    <td class="p-2 w-1/2">{{ swap.BatteryOutSoc.toFixed(2) }} %</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Price:</td>
                                                    <td class="p-2 w-1/2">{{ swap.Cost.toFixed(2) }} RWF</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Distance Covered:</td>
                                                    <td class="p-2 w-1/2">{{ swap.DistanceCovered.toFixed(2) }} Km</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Energy Consumed:</td>
                                                    <td class="p-2 w-1/2">{{ swap.EnergyConsumed.toFixed(2) }} kWh</td>
                                                </tr>
                                                <tr class="border">
                                                    <td class="p-2 w-1/2">Driver Names:</td>
                                                    <td class="p-2 w-1/2">{{ swap.Driver.LastName }} - {{ swap.Driver.FirstName }}</td>
                                                </tr>
                                            </tbody>
                                        </table>
                                    </div>
                                </td>
                            </tr>
                        </template>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
