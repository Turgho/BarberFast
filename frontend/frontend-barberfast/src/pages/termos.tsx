import Navbar from "@/components/navbar";
import "@/styles/globals.css";

export default function Termos() {
    return (
        <>
            <Navbar />
            <section className="bg-mainColor min-h-screen flex flex-col justify-center items-center p-8">
                <div className="max-w-2xl bg-gray-800 p-8 rounded-xl shadow-lg text-white">
                    <h1 className="text-3xl font-bold text-laranja mb-4">Termos e Condições</h1>
                    <p className="mb-4">
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque luctus odio et velit dapibus, 
                        ut pharetra purus dignissim. Suspendisse potenti.
                    </p>
                    <p className="mb-4">
                        Fusce nec nulla at eros facilisis vehicula. Vestibulum suscipit lorem ac purus ornare tincidunt. 
                        Vivamus et metus nec metus pellentesque elementum.
                    </p>
                    <p className="mb-4">
                        Ao se cadastrar, você concorda com os termos acima.
                    </p>
                </div>
            </section>
        </>
    );
}
