const avaliacoes = [
    { text: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", nome: "usuário 1", img: "/images/avaliacao-1.jpg" },
    { text: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", nome: "usuário 2", img: "/images/avaliacao-2.jpg" },
    { text: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.", nome: "usuário 3", img: "/images/avaliacao-3.jpg" }
];

export default function Avaliacoes() {
    return (
        <section id="avaliacoes" className="py-[30%] md:p-[2%] bg-mainColor">
            <h1 className="text-center font-archivo-black text-[40px] p-10">AVALIAÇÕES</h1>
            <div className="avaliacoes grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-[5%] md:px-[5%] h-full lg:h-[600px] justify-center items-center">

                {avaliacoes.map((avaliacao, index) => (
                    <div className="cards p-6 bg-laranja rounded-lg shadow-lg flex flex-col justify-between items-center" key={index}>

                        <div className="flex items-start mb-4 w-full">
                            <img src="/icons/quote.svg" alt="quote-icon" width="60" height="60"/>
                        </div>

                        <p className="text-sm md:text-base lg:text-lg text-center text-white">{avaliacao.text}</p>

                        <div className="usuario flex flex-row justify-around items-center w-full">
                            <h1 className="font-archivo-black text-[20px] md:text-[24px] text-white mb-2">{avaliacao.nome.toUpperCase()}</h1>
                            <div className="div-usuario flex justify-center items-center w-[80px] h-[80px] border-white border-2 rounded-full overflow-hidden mb-4">
                                <img src={avaliacao.img || "/icons/default-avatar.svg"} alt="foto-usuario" className="object-cover w-full h-full" />
                            </div>
                        </div>
                    </div>
                ))}

            </div>
        </section>
    );
}
