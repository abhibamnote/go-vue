import MainLayout from "@/layouts/MainLayout.vue";
import ChartView from "@/views/ChartView.vue";
import IndexView from "@/views/IndexView.vue";
import OptionChain from "@/views/OptionChain.vue";
import StrategyView from "@/views/StrategyView.vue";


const routes = [
    {
        path: "/",
        component: MainLayout,
        children: [
            {
                path: "/",
                name: "Index",
                component: IndexView,
            },
            {
                path: "/chart",
                name: "Chart",
                component: ChartView
            },
            {
                path: "/option-chain",
                name: "Option Chain",
                component: OptionChain
            },
            {
                path: "/strategy",
                name: "strategy",
                component: StrategyView
            }
        ],
    },
];

export default routes;
