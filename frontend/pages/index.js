import Checkout from "../components/index/checkout";
import MainHeader from "../components/index/header";
import InputCard from "../components/index/input";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'


export default function Home() {
  const queryClient = new QueryClient()
  return (
    <QueryClientProvider client={queryClient}>
      <div className="flex flex-row justify-center">
        <div className="min-w-[300px] max-w-[800px] px-5 pb-5">
          <MainHeader className="mt-5" />
          <InputCard className="mt-5" />
          <Checkout className="mt-5" />
        </div>
      </div>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}
