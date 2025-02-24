import TipoServicos from "@/components/tipoServicos";

export default function Servicos(){
    return (
        <section id="servicos">
            <div className="servicos bg-mainColor">
                <h1 className="text-center font-archivo-black text-[40px] p-[5%]">SERVIÇOS</h1>  
                <TipoServicos />
            </div>
        </section>
    );
}