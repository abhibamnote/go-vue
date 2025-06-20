<script setup>
import { onMounted, reactive, computed, ref } from "vue";
import { useAuthStore } from "@/store/authStore";
import axios from "axios";
const auth = useAuthStore();

const dropdown = ref(false);

const state = reactive({
    search: "",
    segment: "",
    masterData: [],
    watchlist: [],
    currentPage: 1,
    itemsPerPage: 10,
    isLoading: false,
    error: null,
    sortBy: "token",
    sortDirection: "asc",
});

const sortedWatchlist = computed(() => {
    return [...state.watchlist].sort((a, b) => {
        let modifier = state.sortDirection === "asc" ? 1 : -1;
        if (a.Master[state.sortBy] < b.Master[state.sortBy])
            return -1 * modifier;
        if (a.Master[state.sortBy] > b.Master[state.sortBy])
            return 1 * modifier;
        return 0;
    });
});

const paginatedData = computed(() => {
    const start = (state.currentPage - 1) * state.itemsPerPage;
    const end = start + state.itemsPerPage;
    return state.masterData.slice(start, end);
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
        console.log(state.masterData);
        
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
const showDropdown = () => {
    dropdown.value = true;
}

const hideDropdown = () => {
    dropdown.value = false;
}

onMounted(fetchMasterData);
</script>

<template>
    <section>
        <v-container>
            <v-row>
                <v-col cols="12" md="6">
                    <v-select
                        v-model="state.segment"
                        label="Segment"
                        :items="['', 'BFO', 'BSE', 'CDS', 'MCX', 'NFO', 'NSE']"
                        clearable
                        @update:modelValue="fetchMasterData"
                    ></v-select>
                </v-col>
                <v-col cols="12" md="6" style="position: relative;">
                    <v-text-field
                        v-model="state.search"
                        label="Search"
                        @focus="showDropdown"
                        @blur="hideDropdown"
                        placeholder="Search by symbol, ISIN, etc."
                        clearable
                        @update:modelValue="fetchMasterData"
                    >
                        <div v-if="dropdown" class="dropdown">
                            <div class="dropdown-item" v-for="item in state.masterData" :key="item.ID">
                                {{ item.tradingSymbol }} - {{ item.segment }}
                                <v-btn
                                    color="primary"
                                    v-if="!checkInWatchList(item.ID)"
                                    @click="addToWatchList(item.ID)"
                                    >Add</v-btn
                                >
                                <!-- <v-btn
                                    color="error"
                                    v-if="checkInWatchList(item.ID)"
                                    @click="deleteFromWatchList(item.ID)"
                                    >Delete</v-btn
                                > -->
                            </div>
                        </div>
                    </v-text-field>
                </v-col>
            </v-row>
        </v-container>

        <v-container>
            <v-alert v-if="state.error" type="error" class="mb-4">
                {{ state.error }}
            </v-alert>
        </v-container>

        <v-container>
            <v-table>
                <thead>
                    <tr>
                        <th>Delete</th>
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
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="item in sortedWatchlist" :key="item.ID">
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
.dropdown {
    position: absolute;
    width: 100%;
    max-height: 150px;
    z-index: 999;
    background-color: white;
    overflow-y: scroll;
    border: 1px solid #ccc;
    top: 100%;
}
.dropdown-show {
    display: block;
}
.dropdown-item {
    padding: 0.5rem;
    border-bottom: 1px solid #ccc;
    display: flex;
    padding: 0.5rem 2rem;
    align-items: center;
}
.dropdown-item button{
    margin-left: 2rem;
}
</style>
