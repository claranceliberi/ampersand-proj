import { TelemetryData, agent, BatteryData } from '../types/index';
<script lang="ts" setup>
import { BatteryData } from "~/types/index"
import { formatTime } from "~/utils/index"

const { data: telematics } = await useFetch<BatteryData[]>("http://localhost:8059/api/v1/battery_telematics/")
</script>

<template>
    <div>
        <div>
            <div class="relative overflow-x-auto">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" class="px-6 py-3">ID</th>
                            <th scope="col" class="px-6 py-3">Lat, Lang</th>
                            <th scope="col" class="px-6 py-3">Battery</th>
                            <th scope="col" class="px-6 py-3">Battery SOC</th>
                            <th scope="col" class="px-6 py-3">Is Charging?</th>
                            <th scope="col" class="px-6 py-3">ChargingRate?</th>
                            <th scope="col" class="px-6 py-3">createdAt</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="telematic in telematics" :key="telematic.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ telematic.ID }}</th>
                            <td class="px-6 py-4">{{ telematic.Latitude }}, {{ telematic.Longitude }}</td>
                            <td class="px-6 py-4">{{ telematic.BatteryID }}</td>
                            <td class="px-6 py-4">{{ telematic.BatterySoc }}</td>
                            <td class="px-6 py-4">{{ telematic.Charging }}</td>
                            <td class="px-6 py-4">{{ telematic.ChargingRate }}</td>
                            <td class="px-6 py-4">{{ formatTime(telematic.CreatedAt) }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
