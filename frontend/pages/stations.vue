import { TelemetryData, station } from '../types/index';
<script lang="ts" setup>
import { Station } from "~/types/index"
import { formatTime } from "~/utils/index"

const { data: stations } = await useFetch<Station[]>("http://localhost:8059/api/v1/stations/")
</script>

<template>
    <div>
        <div>
            <div class="relative overflow-x-auto">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" class="px-6 py-3">ID</th>
                            <th scope="col" class="px-6 py-3">Location</th>
                            <th scope="col" class="px-6 py-3">Description</th>
                            <th scope="col" class="px-6 py-3">Created On</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="station in stations" :key="station.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ station.ID }}</th>
                            <td class="px-6 py-4">{{ station.Location }}</td>
                            <td class="px-6 py-4">{{ station.Description }}</td>
                            <td class="px-6 py-4">{{ formatTime(station.CreatedAt) }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
