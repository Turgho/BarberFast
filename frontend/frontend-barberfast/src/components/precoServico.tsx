const precoServico = [
    { text: "Corte", preco: "R$30" },
    { text: "Barba", preco: "R$30" },
    { text: "Degradê", preco: "R$30" },
    { text: "Hidratação", preco: "R$30" },
    { text: "Sobrancelha", preco: "R$30" },
    { text: "Completo", preco: "R$30" },
];

export default function PrecoServicos() {
    return (
        <div className="precos flex flex-col justify-center items-center gap-4 md:mr-[50%] bg-laranja w-full md:w-[25%] h-[57%] absolute">
            <h1 className="text-white font-archivo-black text-[30px] text-center">
                LISTA DE PREÇOS
            </h1>
            {precoServico.map((servico, index) => (
                <div key={index} className="flex flex-row justify-between w-full px-[10%] items-center text-center text-[25px]">
                    <p className="font-julius-one">{servico.text}</p>
                    <p className="text-black font-archivo-black">{servico.preco}</p>
                </div>
            ))} 
        </div>
    );
}