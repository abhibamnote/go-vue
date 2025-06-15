<script setup>
import { useAuthStore } from "@/store/authStore";
import axios from "axios";
import { onMounted, reactive, watch } from "vue";
const auth = useAuthStore();

const data = reactive({
    optionData: [],
    aimStrike: 0,
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
    const response = await axios.get(
        import.meta.env.VITE_BACKEND_URL + "/api/option-chain",
        {
            headers: {
                Authorization: auth.token,
            },
        }
    );

    console.log(response);
    data.optionData = response.data.data.optionChainData;
    data.aimStrike = response.data.data.aimStrike;
});

watch(
    () => data.toggles,
    (newToggles) => {
        const activeCount = Object.values(newToggles).filter(Boolean).length;
        data.span = 6 + activeCount;
    },
    { deep: true }
);

const checkHighLight = (strikePrice, type) => {
    if (strikePrice < data.aimStrike && type === "call") {
        return "nhighlight";
    } else if (strikePrice > data.aimStrike && type === "put") {
        return "nhighlight";
    }
};
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

        <v-container style="overflow-x: auto">
            <v-table class="my-table" fixed-header height="600px">
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
                        <th>OI Chng (lots)</th>
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
                        <th>OI Chng (lots)</th>
                        <th v-show="data.toggles.gamma">Gamma</th>
                        <th v-show="data.toggles.vega">Vega</th>
                        <th v-show="data.toggles.theta">Theta</th>
                        <th v-show="data.toggles.delta">Delta</th>
                        <th v-show="data.toggles.imtProb">ITM Prob.</th>
                    </tr>
                </thead>

                <tbody>
                    <tr
                        v-for="row in data.optionData"
                        :key="row.strikePrice"
                        :class="`${
                            data.aimStrike === row.strikePrice
                                ? 'highlight'
                                : ''
                        }`"
                    >
                        <!-- Calls -->
                        <td
                            :class="checkHighLight(row.strikePrice, 'call')"
                            v-show="data.toggles.imtProb"
                        >
                            {{ row.citmp }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'call')"
                            v-show="data.toggles.delta"
                        >
                            {{ row.callDelta }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'call')"
                            v-show="data.toggles.theta"
                        >
                            {{ row.callTheta }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'call')"
                            v-show="data.toggles.vega"
                        >
                            {{ row.callVega }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'call')"
                            v-show="data.toggles.gamma"
                        >
                            {{ row.callGamma }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'call')">
                            0
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'call')">
                            {{ row.callOi }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'call')">
                            {{ row.callVolume }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'call')">
                            {{ row.callIV }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'call')">
                            {{ row.callLtp }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'call')">
                            0
                        </td>

                        <td style="font-weight: 600">
                            {{ row.strikePrice }}
                        </td>

                        <!-- Puts -->
                        <td :class="checkHighLight(row.strikePrice, 'put')">
                            0
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'put')">
                            {{ row.putLtp }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'put')">
                            {{ row.putIv }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'put')">
                            {{ row.putVolume }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'put')">
                            {{ row.putOi }}
                        </td>
                        <td :class="checkHighLight(row.strikePrice, 'put')">
                            0
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'put')"
                            v-show="data.toggles.gamma"
                        >
                            {{ row.putGamma }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'put')"
                            v-show="data.toggles.vega"
                        >
                            {{ row.putVega }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'put')"
                            v-show="data.toggles.theta"
                        >
                            {{ row.putTheta }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'put')"
                            v-show="data.toggles.delta"
                        >
                            {{ row.putDelta }}
                        </td>
                        <td
                            :class="checkHighLight(row.strikePrice, 'put')"
                            v-show="data.toggles.imtProb"
                        >
                            {{ row.pitmp }}
                        </td>
                    </tr>
                </tbody>
            </v-table>
        </v-container>
    </section>
</template>

<style scoped>
.my-table table {
    border-collapse: collapse;
    width: 100%;
}
.my-table th,
.my-table td {
    border: 1px solid black;
    min-width: 75px;
    font-size: 13px;
    font-weight: 600;
}
.my-table th {
    background: white;
    position: sticky;
    top: 0;
    z-index: 1;
    color: #4682b4;
    text-align: center;
}
.my-table td {
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
.highlight {
    background-color: #cefc8a !important;
}
.nhighlight {
    background-color: #f1eed9 !important;
}
</style>
