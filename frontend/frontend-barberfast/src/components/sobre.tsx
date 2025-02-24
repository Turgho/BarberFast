export default function Sobre(){
    return (
        <section id="sobre">
            <div className="sobre flex h-[600px] bg-home-2 bg-cover bg-fixed relative">
                <div className="overlay flex flex-col justify-center items-center gap-[15%] text-center h-full bg-mainColor/75 p-16 font-julius sm:text-[20px] md:text-[25px] absolute">
                    <h1 className="text-[35px] font-archivo-black">
                        SOBRE NÓS!
                    </h1>
                    <p className="px-[10%]">
                        Na BarberFast, acreditamos que um corte de cabelo ou qualquer outro serviço de
                        barbearia de qualidade não precisa ser uma experiência demorada.
                        Nosso objetivo é oferecer um atendimento rápido, eficiente e de excelência,
                        para que nossos clientes possam aproveitar o melhor da beleza e do cuidado com o estilo, sem perder tempo.
                    </p>
                </div>
            </div>
        </section>
    );
}