import MainLayout from "@/layouts/MainLayout.vue";
import ChartView from "@/views/ChartView.vue";
import IndexView from "@/views/IndexView.vue";
import OptionChain from "@/views/OptionChain.vue";


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
            }
        ],
    },
];

export default routes;
