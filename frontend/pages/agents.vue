import { TelemetryData, agent } from '../types/index';
<script lang="ts" setup>
import { Agent } from "~/types/index"
import { formatTime } from "~/utils/index"

const { data: agents } = await useFetch<Agent[]>("http://localhost:8059/api/v1/agents/")
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
                        <tr v-for="agent in agents" :key="agent.ID" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                            <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ agent.ID }}</th>
                            <td class="px-6 py-4">{{ agent.FirstName }} {{ agent.LastName }}</td>
                            <td class="px-6 py-4">{{ agent.Email }}</td>
                            <td class="px-6 py-4">{{ agent.PhoneNumber }}</td>
                            <td class="px-6 py-4">{{ formatTime(agent.CreatedAt) }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>
