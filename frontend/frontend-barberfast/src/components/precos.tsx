import PrecoServicos from "./precoServico";

export default function Precos(){
    return (
        <section id="precos">
            <div className="servico-preco flex justify-center items-center bg-mainColor h-[900px]">
                <PrecoServicos/>
                <div className="preco-image bg-precos bg-cover h-[80%] w-full ml-[32%] invisible md:visible"></div>
            </div>
        </section>
    );
}