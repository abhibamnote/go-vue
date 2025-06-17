<script setup>
import { onMounted, reactive, computed } from "vue";
import { useAuthStore } from "@/store/authStore";
import axios from "axios";
const auth = useAuthStore();

const state = reactive({
    search: "",
    segment: "",
    masterData: [],
    watchlist: [],
    currentPage: 1,
    itemsPerPage: 10,
    isLoading: false,
    error: null,
    sortBy: "segment",
    sortDirection: "asc",
});

const sortedData = computed(() => {
    return [...state.masterData].sort((a, b) => {
        let modifier = state.sortDirection === "asc" ? 1 : -1;
        if (a[state.sortBy] < b[state.sortBy]) return -1 * modifier;
        if (a[state.sortBy] > b[state.sortBy]) return 1 * modifier;
        return 0;
    });
});

const paginatedData = computed(() => {
    const start = (state.currentPage - 1) * state.itemsPerPage;
    const end = start + state.itemsPerPage;
    return sortedData.value.slice(start, end);
});

const totalPages = computed(() => {
    return Math.ceil(state.masterData.length / state.itemsPerPage);
});

const sortTable = (column) => {
    if (state.sortBy === column) {
        state.sortDirection = state.sortDirection === "asc" ? "desc" : "asc";
    } else {
        state.sortBy = column;
        state.sortDirection = "asc";
    }
    state.currentPage = 1;
};

const fetchMasterData = async () => {
    try {
        state.isLoading = true;
        state.error = null;
        const response = await axios.get(
            import.meta.env.VITE_BACKEND_URL + "/api/master-data",
            {
                params: {
                    segment: state.segment,
                    search: state.search.toUpperCase(),
                },
                headers: {
                    Authorization: auth.token,
                },
            }
        );
        state.masterData = response.data.data.masterData;
        state.currentPage = 1;

        const responseTwo = await axios.get(
            import.meta.env.VITE_BACKEND_URL + "/api/watchlist",
            {
                headers: {
                    Authorization: auth.token,
                },
            }
        );
        state.watchlist = responseTwo.data.data.watchList;
    } catch (err) {
        console.error("Error fetching master data:", err);
        state.error = "Failed to load data. Please try again later.";
    } finally {
        state.isLoading = false;
    }
};

const addToWatchList = async (id) => {
    const response = await axios.post(
        import.meta.env.VITE_BACKEND_URL + "/api/watchlist",
        { masterId: id },
        { headers: { Authorization: auth.token } }
    );
    state.watchlist.push(response.data.data.watchList);
};

const deleteFromWatchList = async (id) => {
    await axios.delete(
        import.meta.env.VITE_BACKEND_URL + "/api/watchlist/" + id,
        { headers: { Authorization: auth.token } }
    );
    state.watchlist = state.watchlist.filter((item) => item.ID !== id);
};

const checkInWatchList = (id) => {
    return state.watchlist.some((item) => item.Master.ID === id);
};

onMounted(fetchMasterData);
</script>

<template>
    <section>
        <v-container>
            <v-row>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="state.search"
                        label="Search"
                        placeholder="Search by symbol, ISIN, etc."
                        clearable
                        @update:modelValue="fetchMasterData"
                    ></v-text-field>
                </v-col>
                <v-col cols="12" md="6">
                    <v-select
                        v-model="state.segment"
                        label="Segment"
                        :items="['', 'BFO', 'BSE', 'CDS', 'MCX', 'NFO', 'NSE']"
                        clearable
                        @update:modelValue="fetchMasterData"
                    ></v-select>
                </v-col>
            </v-row>
        </v-container>

        <v-container>
            <v-alert v-if="state.error" type="error" class="mb-4">
                {{ state.error }}
            </v-alert>

            <v-card v-if="state.isLoading">
                <v-card-text class="text-center py-8">
                    <v-progress-circular
                        indeterminate
                        color="primary"
                    ></v-progress-circular>
                    <div class="mt-4">Loading master data...</div>
                </v-card-text>
            </v-card>

            <template v-else>
                <v-table fixed-header height="600px">
                    <thead>
                        <tr>
                            <th>Add</th>
                            <th @click="sortTable('segment')">
                                Segment
                                <font-awesome-icon
                                    v-if="state.sortBy === 'segment'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('token')">
                                Token
                                <font-awesome-icon
                                    v-if="state.sortBy === 'token'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('symbol')">
                                Symbol
                                <font-awesome-icon
                                    v-if="state.sortBy === 'symbol'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('tradingSymbol')">
                                Trading Symbol
                                <font-awesome-icon
                                    v-if="state.sortBy === 'tradingSymbol'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('instrumentType')">
                                Intrument Type
                                <font-awesome-icon
                                    v-if="state.sortBy === 'instrumentType'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('expiryDate')">
                                Expiry Date
                                <font-awesome-icon
                                    v-if="state.sortBy === 'expiryDate'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('tickSize')">
                                Tick Size
                                <font-awesome-icon
                                    v-if="state.sortBy === 'tickSize'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('lotSize')">
                                Lot Size
                                <font-awesome-icon
                                    v-if="state.sortBy === 'lotSize'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('optionType')">
                                Option Type
                                <font-awesome-icon
                                    v-if="state.sortBy === 'optionType'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('strike')">
                                Strike
                                <font-awesome-icon
                                    v-if="state.sortBy === 'strike'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('pricePrec')">
                                Price Precision
                                <font-awesome-icon
                                    v-if="state.sortBy === 'pricePrec'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('multiplier')">
                                Multiplier
                                <font-awesome-icon
                                    v-if="state.sortBy === 'multiplier'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('isin')">
                                Isin
                                <font-awesome-icon
                                    v-if="state.sortBy === 'isin'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('priceMultiplier')">
                                Price Multiplier
                                <font-awesome-icon
                                    v-if="state.sortBy === 'priceMultiplier'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                            <th @click="sortTable('companyName')">
                                Company Name
                                <font-awesome-icon
                                    v-if="state.sortBy === 'companyName'"
                                    :icon="
                                        state.sortDirection === 'asc'
                                            ? ['fas', 'arrow-up']
                                            : ['fas', 'arrow-down']
                                    "
                                />
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="item in paginatedData" :key="item.token">
                            <td>
                                <v-btn
                                    v-if="!checkInWatchList(item.ID)"
                                    color="primary"
                                    @click="addToWatchList(item.ID)"
                                    >Add</v-btn
                                >
                                <v-btn v-else color="success">Added</v-btn>
                            </td>
                            <td>{{ item.segment }}</td>
                            <td>{{ item.token }}</td>
                            <td>{{ item.symbol }}</td>
                            <td>{{ item.tradingSymbol }}</td>
                            <td>{{ item.instrumentType }}</td>
                            <td>{{ item.expiryDate }}</td>
                            <td>{{ item.tickSize }}</td>
                            <td>{{ item.lotSize }}</td>
                            <td>{{ item.optionType }}</td>
                            <td>{{ item.strike }}</td>
                            <td>{{ item.pricePrec }}</td>
                            <td>{{ item.multiplier }}</td>
                            <td>{{ item.isin }}</td>
                            <td>{{ item.priceMultiplier }}</td>
                            <td>{{ item.companyName }}</td>
                        </tr>
                    </tbody>
                </v-table>

                <div class="text-center mt-4">
                    <v-pagination
                        v-model="state.currentPage"
                        :length="totalPages"
                        :total-visible="7"
                        rounded
                    ></v-pagination>
                    <div class="d-flex align-center justify-center mt-2">
                        <span class="mr-2">Items per page:</span>
                        <v-select
                            v-model="state.itemsPerPage"
                            :items="[10, 25, 50, 100]"
                            density="compact"
                            style="max-width: 100px"
                            @update:modelValue="state.currentPage = 1"
                        ></v-select>
                        <span class="ml-2">
                            {{
                                (state.currentPage - 1) * state.itemsPerPage + 1
                            }}
                            -
                            {{
                                Math.min(
                                    state.currentPage * state.itemsPerPage,
                                    state.masterData.length
                                )
                            }}
                            of {{ state.masterData.length }} items
                        </span>
                    </div>
                </div>
                <v-alert
                    v-if="!state.masterData.length"
                    type="info"
                    class="mt-4"
                    >No data found.</v-alert
                >
            </template>
        </v-container>

        <v-container>
            <v-table>
                <thead>
                    <tr>
                        <th>Delete</th>
                        <th>Token</th>
                        <th>Symbol</th>
                        <th>Trading Symbol</th>
                        <th>Segment</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in state.watchlist" :key="item.ID">
                        <td>
                            <v-btn
                                color="error"
                                @click="deleteFromWatchList(item.ID)"
                                >Delete</v-btn
                            >
                        </td>
                        <td>{{ item.Master.token }}</td>
                        <td>{{ item.Master.symbol }}</td>
                        <td>{{ item.Master.tradingSymbol }}</td>
                        <td>{{ item.Master.segment }}</td>
                    </tr>
                </tbody>
            </v-table>
        </v-container>
    </section>
</template>

<style scoped>
.v-table {
    width: 100%;
}
.v-table td {
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    max-width: 200px;
}
.v-table th {
    cursor: pointer;
    user-select: none;
}
.v-table th:hover {
    background-color: rgba(0, 0, 0, 0.04);
}
.font-awesome-icon {
    font-size: 16px;
    margin-left: 4px;
    vertical-align: middle;
}
</style>
