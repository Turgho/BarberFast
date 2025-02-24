const fotos = [
    { image: '/images/galeria-1.jpg', nome: 'João', rede: 'https://www.instagram.com/' },
    { image: '/images/galeria-2.jpg', nome: 'Marcos', rede: 'https://www.instagram.com/' },
    { image: '/images/galeria-3.jpg', nome: 'Matheus', rede: 'https://www.instagram.com/' },
    { image: '/images/galeria-4.jpg', nome: 'Artur', rede: 'https://www.instagram.com/' },
    { image: '/images/galeria-5.jpg', nome: 'Caio', rede: 'https://www.instagram.com/' },
    { image: '/images/galeria-6.jpg', nome: 'Pedro', rede: 'https://www.instagram.com/' },
    { image: '/images/galeria-7.jpg', nome: 'Gabriel', rede: 'https://www.instagram.com/' },
    { image: '/images/galeria-8.jpg', nome: 'Jhonas', rede: 'https://www.instagram.com/' }
];

export default function Galeria() {
    return (
        <section id="galeria" className="py-[5%] bg-mainColor">
            <h1 className="text-center font-archivo-black text-[40px] p-10">GALERIA</h1>
            <div className="galeria w-full grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 border-y-4 border-laranja">
                {fotos.map((foto, index) => (
                    <div className="foto relative w-full h-[300px] cursor-pointer overflow-hidden" key={index}>
                        {/* Imagem */}
                        <img 
                            src={foto.image} 
                            alt={`Foto de ${foto.nome}`} 
                            className="w-full h-full object-cover"
                        />

                        {/* Overlay */}
                        <div className="absolute top-0 left-0 w-full h-full bg-mainColor/50 opacity-0 hover:opacity-100 transition-opacity duration-300 flex flex-col items-center justify-center">
                            <a href={foto.rede} target="_blank" rel="noopener noreferrer" className="flex flex-col items-center">
                                <span className="text-white text-[30px] font-bold">{foto.nome.toUpperCase()}</span>
                                <img src="/icons/footer/instagram.svg" alt="Ícone do Instagram" width="50px" />
                            </a>
                        </div>
                    </div>
                ))}
            </div>
        </section>
    );
}

