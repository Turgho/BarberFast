import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import { format } from 'date-fns';
import { toZonedTime } from 'date-fns-tz';
import Navbar from '@/components/navbar';
import "@/styles/globals.css";

interface FormData {
    usuario_id: string;
    servico_id: number;
    data_inicio: string;
    data_fim: string;
    status: string;
}

interface Servico {
    id: number;
    nome: string;
    preco: number;
    duracao_minutos: number;
}

interface Errors {
    [key: string]: string;
}

export default function Agendar() {
    useEffect(() => {
        const token = sessionStorage.getItem("token");
        if (!token) {
            router.push("/login"); // Redireciona se não houver token
        }
    }, []);

    const router = useRouter();

    const [formData, setFormData] = useState<FormData>({
        usuario_id: "",
        servico_id: 0,
        data_inicio: "",
        data_fim: "",
        status: "confirmado",
    });

    const [servicos, setServicos] = useState<Servico[]>([]);
    const [errors, setErrors] = useState<Errors>({});
    const [isSuccess, setIsSuccess] = useState<boolean>(false);
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
    const [servicoSelecionado, setServicoSelecionado] = useState<Servico | null>(null);

    useEffect(() => {
        const userId = sessionStorage.getItem("user_id");
        if (userId) {
            setFormData(prevData => ({ ...prevData, usuario_id: userId }));
        }

        async function fetchData() {
            try {
                const response = await fetch("http://192.168.1.4:5050/v1/usuario/servicos", {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + sessionStorage.getItem("token"),
                    },
                });
                if (response.ok) {
                    const servicosData: Servico[] = await response.json();
                    setServicos(servicosData);
                }
            } catch (error) {
                console.error("Erro ao buscar serviços:", error);
            }
        }
        fetchData();
    }, []);

    const calcularDataFim = (dataInicio: string, duracao: number) => {
        const dataInicioObj = new Date(dataInicio);
        return new Date(dataInicioObj.getTime() + duracao * 60000).toISOString().slice(0, 16);
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
    
        setFormData((prevData) => ({
            ...prevData,
            [name]: name === "servico_id" ? Number(value) : value,
        }));
    
        if (name === "servico_id") {
            const servico = servicos.find(s => s.id === Number(value));
            setServicoSelecionado(servico || null);
    
            if (servico && formData.data_inicio) {
                const dataFim = calcularDataFim(formData.data_inicio, servico.duracao_minutos);
                setFormData((prevData) => ({
                    ...prevData,
                    data_fim: dataFim,
                }));
            }
        }
    
        if (name === "data_inicio") {
            const servico = servicos.find(s => s.id === formData.servico_id);
            if (servico) {
                const dataFim = calcularDataFim(value, servico.duracao_minutos);
                setFormData((prevData) => ({
                    ...prevData,
                    data_fim: dataFim,
                }));
            }
        }
    };
    
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
    
        let validationErrors: Errors = {};
        setIsSubmitting(true);
    
        if (!formData.usuario_id) validationErrors.usuario_id = "O cliente é obrigatório.";
        if (!formData.servico_id) validationErrors.servico_id = "O serviço é obrigatório.";
        if (!formData.data_inicio) validationErrors.data_inicio = "A data de início é obrigatória.";
    
        const fusoHorario = "America/Sao_Paulo";
    
        // Converter data_inicio para UTC e formatar corretamente
        const dataInicioUtc = toZonedTime(new Date(formData.data_inicio), fusoHorario);
        const dataInicioFormatada = format(dataInicioUtc, "yyyy-MM-dd'T'HH:mm:ssXXX");
    
        const servicoSelecionado = servicos.find(s => s.id === Number(formData.servico_id));
        if (servicoSelecionado) {
            const dataFimCalculada = new Date(dataInicioUtc.getTime() + servicoSelecionado.duracao_minutos * 60000);
            const dataFimFormatada = format(toZonedTime(dataFimCalculada, fusoHorario), "yyyy-MM-dd'T'HH:mm:ssXXX");
    
            try {
                const response = await fetch("http://192.168.1.4:5050/v1/usuario/agendar", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "Authorization": "Bearer " + sessionStorage.getItem("token"),
                    },
                    body: JSON.stringify({
                        ...formData,
                        data_inicio: dataInicioFormatada,
                        data_fim: dataFimFormatada,
                    }),
                });
    
                if (response.ok) {
                    setIsSuccess(true);
                    setTimeout(() => {
                        router.push("/")
                    }, 3000);
                } else {
                    const data = await response.json();
                    setErrors({ server: data.message || "Erro desconhecido" });
                }
            } catch (error) {
                console.error("Erro ao agendar:", error);
                setErrors({ server: "Erro na conexão com o servidor" });
            } finally {
                setIsSubmitting(false);
            }
        } else {
            setErrors(validationErrors);
            setIsSubmitting(false);
        }
    };
    
    return (
        <>
            <Navbar/>
            <section id="agendamento" className="bg-cadastro bg-cover min-h-screen flex flex-col justify-center items-center">

                <div className="overlay flex flex-col justify-center items-center bg-mainColor/75 absolute w-full h-full">

                    <div className="agendamento flex flex-col md:flex-row items-center justify-center gap-8 p-6 md:p-12 mt-16">

                        <div className="box-esquerda flex-1 text-center md:text-left px-4 md:px-[10%]">
                            <h1 className="text-xl text-center md:text-3xl text-white mb-6">
                            Agende seu <span className="text-white">SERVIÇO</span> com facilidade.
                            </h1>
                        </div>

                        <div className="box-direita flex-1 flex items-center justify-center p-6 rounded-xl w-full max-w-md">
                            <form onSubmit={handleSubmit} className="w-full font-archivo">
                                <h2 className="text-white text-2xl font-archivo-black mb-6 text-center">AGENDAR SERVIÇO</h2>
                                
                                <div className="mb-4">
                                    <input
                                        type="hidden"
                                        name="usuario_id"
                                        value={formData.usuario_id}
                                        onChange={handleChange}
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-white"
                                        disabled
                                    />
                                </div>

                                <div className="mb-4">
                                    <select
                                        name="servico_id"
                                        value={formData.servico_id}
                                        onChange={handleChange}
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    >
                                        <option value="">Selecione um serviço</option>
                                        {servicos.map(servico => (
                                            <option key={servico.id} value={servico.id}>
                                                {servico.nome}
                                            </option>
                                        ))}
                                    </select>
                                    {errors.servico_id && <span className="text-red-500">{errors.servico_id}</span>}
                                </div>

                                {servicoSelecionado && (
                                    <div className="text-white text-sm mb-4">
                                        <p>Serviço: {servicoSelecionado.nome}</p>
                                        <p>Preço: R${servicoSelecionado.preco}</p>
                                        <p>Duração: {servicoSelecionado.duracao_minutos} minutos</p>
                                    </div>
                                )}

                                <div className="mb-4">
                                    <label htmlFor="data_inicio" className="text-white">Data de Início</label>
                                    <input
                                        type="datetime-local"
                                        id="data_inicio"
                                        name="data_inicio"
                                        value={formData.data_inicio}
                                        onChange={handleChange}
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-white"
                                        required
                                    />
                                    {errors.data_inicio && <span className="text-red-500">{errors.data_inicio}</span>}
                                </div>

                                <div className="mb-4">
                                    <input
                                        type="hidden"
                                        id="data_fim"
                                        name="data_fim"
                                        value={formData.data_fim}
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-white"
                                        disabled
                                    />
                                    {errors.data_fim && <span className="text-red-500">{errors.data_fim}</span>}
                                </div>

                                <div className="flex justify-center items-center mt-8">
                                    <button
                                        type="submit"
                                        disabled={isSubmitting}
                                        className="px-6 py-2 rounded-full bg-laranja text-white font-bold hover:bg-laranja-dark focus:outline-none"
                                    >
                                        {isSubmitting ? "Agendando..." : "Agendar"}
                                    </button>
                                </div>
                            </form>
                        </div>
                    </div>
                    {/* Mensagem de sucesso */}
                    {isSuccess && (
                        <div className="fixed bottom-10 left-0 right-0 bg-green-500 text-white text-center py-3">
                            Agendamento realizado com sucesso! Redirecionando...
                        </div>
                    )}
                </div>
            </section>
        </>
    );
}
