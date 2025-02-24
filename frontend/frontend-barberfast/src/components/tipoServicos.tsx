import React from 'react';

// Array de objetos que contém informações sobre os serviços oferecidos
const servicos = [
    {
        tipo: "corte",
        text: "Serviço que inclui o corte de cabelo conforme o estilo preferido do cliente.",
        icon: "/icons/scissor.svg"
    },
    {
        tipo: "barba",
        text: "Serviço de modelagem de barba, com uso de navalha e produtos específicos para um acabamento perfeito.",
        icon: "/icons/beard.svg"
    },
    {
        tipo: "degradê",
        text: "Corte de cabelo com transição suave entre diferentes comprimentos de cabelo, criando um efeito de gradação ou degradê.",
        icon: "/icons/razor.svg"
    },
    {
        tipo: "hidratação",
        text: "Tratamento capilar que promove hidratação e saúde aos fios, ideal para cabelos secos ou danificados.",
        icon: "/icons/shampoo.svg"
    },
    {
        tipo: "sobrancelhas",
        text: "Design de sobrancelhas, utilizando técnicas de depilação para deixar o rosto mais harmônico.",
        icon: "/icons/eyebrow.svg"
    },
    {
        tipo: "completo",
        text: "Pacote completo com corte + barba + sobrancelha, para manter o visual em dia.",
        icon: "/icons/premium.svg"
    }
];

export default function TipoServicos() {
    return (
        // Container principal com grid responsivo
        <div className="servicos grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-10 py-10 px-5 md:px-10 text-white">
            {servicos.map((servico, index) => (
                // Card de cada serviço com efeito de hover (escala)
                <div 
                    className="cards flex flex-col justify-center items-center text-center space-y-4 transition-transform duration-300 hover:scale-105"
                    key={index}
                >
                    {/* Container do ícone do serviço */}
                    <div className="back flex justify-center items-center bg-laranja rounded-full w-[120px] h-[120px] md:w-[150px] md:h-[150px] shadow-lg">
                        {/* Ícone com efeito de hover para escalar */}
                        <img 
                            src={servico.icon} 
                            alt={servico.tipo} 
                            className="w-[70px] md:w-[100px] transition-transform duration-300 hover:scale-110"
                        />
                    </div>
                    {/* Título do serviço */}
                    <h1 className="font-archivo-black text-xl md:text-2xl">
                        {servico.tipo.toUpperCase()}
                    </h1>
                    {/* Descrição do serviço com limite de largura para melhor legibilidade */}
                    <p className="text-sm md:text-base max-w-[300px]">
                        {servico.text}
                    </p>
                </div>
            ))}
        </div>
    );
}
