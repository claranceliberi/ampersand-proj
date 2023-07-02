import { TelemetryData, agent } from '../types/index';
<script lang="ts" setup>
import { Vehicle } from "~/types/index"
import { formatTime } from "~/utils/index"

const { data: vehicles } = await useFetch<Vehicle[]>("http://localhost:8059/api/v1/vehicles/")
</script>

<template>
    <div>
        <div>
            <div class="relative overflow-x-auto">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" class="px-6 py-3">ID</th>
                            <th scope="col" class="px-6 py-3">Identifier</th>
                            <th scope="col" class="px-6 py-3">Model Name</th>
                            <th scope="col" class="px-6 py-3">Manufacture Date</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="vehicle in vehicles" :key="vehicle.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ vehicle.ID }}</th>
                            <td class="px-6 py-4">{{ vehicle.Identifier }}</td>
                            <td class="px-6 py-4">{{ vehicle.ModelName }}</td>
                            <td class="px-6 py-4">{{ formatTime(vehicle.ManufactureDate) }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
