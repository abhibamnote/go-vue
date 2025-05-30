<script setup>
import { useAuthStore } from "@/store/authStore";
import axios from "axios";
import { onMounted, reactive, watch } from "vue";
const auth = useAuthStore();

const data = reactive({
    optionData: [],
    toggles: {
        gamma: false,
        delta: false,
        theta: false,
        vega: false,
        imtProb: false,
    },
    span: 6,
});

onMounted(async () => {
    const response = await axios.get("http://127.0.0.1:8000/api/option-chain", {
        headers: {
            Authorization: auth.token,
        },
    });

    console.log(response);
    data.optionData = response.data.data.optionChainData;
});

watch(() => data.toggles,(newToggles) => {
        const activeCount = Object.values(newToggles).filter(Boolean).length;
        data.span = 6 + activeCount;
    },
    { deep: true }
);
</script>

<template>
    <section>
        <v-container>
            <div class="btn-group">
                <button
                    v-for="(active, key) in data.toggles"
                    :key="key"
                    @click="data.toggles[key] = !data.toggles[key]"
                    :class="{ selected: active }"
                >
                    {{ key }}
                </button>
            </div>
        </v-container>
        <v-container style="overflow-x: scroll">
            <table>
                <thead>
                    <tr>
                        <th :colspan="data.span">CALLS</th>
                        <th rowspan="2">Strike</th>
                        <th :colspan="data.span">PUTS</th>
                    </tr>
                    <tr>
                        <th v-show="data.toggles.imtProb">ITM Prob.</th>
                        <th v-show="data.toggles.delta">Delta</th>
                        <th v-show="data.toggles.theta">Theta</th>
                        <th v-show="data.toggles.vega">Vega</th>
                        <th v-show="data.toggles.gamma">Gamma</th>
                        <th>
                            OI Chng <br />
                            (lots)
                        </th>
                        <th>OI (lots)</th>
                        <th>Volume (lots)</th>
                        <th>IV</th>
                        <th>LTP</th>
                        <th>Change</th>
                        <th>Change</th>
                        <th>LTP</th>
                        <th>IV</th>
                        <th>Volume (lots)</th>
                        <th>OI (lots)</th>
                        <th>
                            OI Chng <br />
                            (lots)
                        </th>
                        <th v-show="data.toggles.gamma">Gamma</th>
                        <th v-show="data.toggles.vega">Vega</th>
                        <th v-show="data.toggles.theta">Theta</th>
                        <th v-show="data.toggles.delta">Delta</th>
                        <th v-show="data.toggles.imtProb">ITM Prob.</th>
                    </tr>
                </thead>
                <tbody style="height: 800px; overflow-y: scroll">
                    <template v-for="row in data.optionData">
                        <tr>
                            <!-- Calls  -->
                            <td v-show="data.toggles.imtProb">
                                {{ row.citmp }}
                            </td>
                            <td v-show="data.toggles.delta">
                                {{ row.callDelta }}
                            </td>
                            <td v-show="data.toggles.theta">
                                {{ row.callTheta }}
                            </td>
                            <td v-show="data.toggles.vega">
                                {{ row.callVega }}
                            </td>
                            <td v-show="data.toggles.gamma">
                                {{ row.callGamma }}
                            </td>
                            <td>0</td>
                            <td>{{ row.callOi }}</td>
                            <td>{{ row.callVolume }}</td>
                            <td>{{ row.callIV }}</td>
                            <td>{{ row.callLtp }}</td>
                            <td>0</td>

                            <td style="font-weight: 600">
                                {{ row.strikePrice }}
                            </td>

                            <!-- Puts -->
                            <td>0</td>
                            <td>{{ row.putLtp }}</td>
                            <td>{{ row.putIv }}</td>
                            <td>{{ row.putVolume }}</td>
                            <td>{{ row.putOi }}</td>
                            <td>0</td>
                            <td v-show="data.toggles.gamma">
                                {{ row.putGamma }}
                            </td>
                            <td v-show="data.toggles.vega">
                                {{ row.putVega }}
                            </td>
                            <td v-show="data.toggles.theta">
                                {{ row.putTheta }}
                            </td>
                            <td v-show="data.toggles.delta">
                                {{ row.putDelta }}
                            </td>
                            <td v-show="data.toggles.imtProb">
                                {{ row.pitmp }}
                            </td>
                        </tr>
                    </template>
                </tbody>
            </table>
        </v-container>
    </section>
</template>

<style scoped>
table{
    table-layout: auto;
    width: 100%;
}
table,
th,
td {
    border: 1px solid black;
    border-collapse: collapse;
    min-width: 75px;
    font-size: 13px;
    font-weight: 600;
}
th {
    color: #4682b4;
}
td {
    text-align: center;
    font-weight: 400;
    padding: 0.25rem 0.5rem;
}

.btn-group {
    display: flex;
}
button {
    padding: 0.5rem 1rem;
    border: 1px solid #bbb;
    background-color: white;
    cursor: pointer;
    border-radius: 0.25rem;
}
button.selected {
    background-color: #1976d2;
    color: white;
    border-color: #1976d2;
}
</style>
