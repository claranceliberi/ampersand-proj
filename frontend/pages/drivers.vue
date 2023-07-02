import { TelemetryData, driver } from '../types/index';
<script lang="ts" setup>
import { Driver } from "~/types/index"
import { formatTime } from "~/utils/index"

const { data: drivers } = await useFetch<Driver[]>("http://localhost:8059/api/v1/drivers/")
</script>

<template>
    <div>
        <div>
            <div class="relative overflow-x-auto">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                        <tr>
                            <th scope="col" class="px-6 py-3">ID</th>
                            <th scope="col" class="px-6 py-3">Full Name</th>
                            <th scope="col" class="px-6 py-3">Email</th>
                            <th scope="col" class="px-6 py-3">Phone</th>
                            <th scope="col" class="px-6 py-3">Joined On</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="driver in drivers" :key="driver.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ driver.ID }}</th>
                            <td class="px-6 py-4">{{ driver.FirstName }} {{ driver.LastName }}</td>
                            <td class="px-6 py-4">{{ driver.Email }}</td>
                            <td class="px-6 py-4">{{ driver.PhoneNumber }}</td>
                            <td class="px-6 py-4">{{ formatTime(driver.CreatedAt) }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
