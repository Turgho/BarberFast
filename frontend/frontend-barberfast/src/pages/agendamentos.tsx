import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import "@/styles/globals.css";

interface Agendamento {
    id: number;
    servico: {
        nome: string;
        descricao: string;
        preco: number;
        duracao_minutos: number;
    };
    data_inicio: string;
    data_fim: string;
    status: string;
}

export default function Agendamentos() {
    const router = useRouter();
    const [agendamentos, setAgendamentos] = useState<Agendamento[]>([]);
    const [filtro, setFiltro] = useState({ status: "", ordenacao: "recente" });
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        const token = sessionStorage.getItem("token");
        if (!token) {
            router.push("/login"); // Redireciona se não houver token
        } else {
            fetchData();
        }
    }, []);

    async function fetchData() {
        setLoading(true);
        try {
            const queryParams = new URLSearchParams();
            if (filtro.status) queryParams.append("status", filtro.status);
            queryParams.append("ordenacao", filtro.ordenacao);

            const response = await fetch(`http://192.168.1.4:5050/v1/usuario/agendamentos?${queryParams.toString()}`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": "Bearer " + sessionStorage.getItem("token"),
                },
                body: JSON.stringify({
                    id: sessionStorage.getItem("user_id"),
                }),
            });

            if (response.ok) {
                let data: Agendamento[] = await response.json();
                setAgendamentos(data);
            }
        } catch (error) {
            console.error("Erro ao buscar agendamentos:", error);
        } finally {
            setLoading(false);
        }
    }

    return (
        <section className="bg-home-1 bg-cover min-h-screen flex items-center justify-center p-5 relative">
            {/* Overlay aprimorado com blur */}
            <div className="absolute inset-0 bg-black bg-opacity-40 backdrop-blur-md"></div>

            <div className="relative z-10 max-w-3xl w-full bg-white/90 p-6 rounded-2xl shadow-xl">
                <h1 className="text-2xl font-bold text-center text-gray-800 mb-6">
                    Meus Agendamentos
                </h1>

                {/* Filtros */}
                <div className="flex flex-col sm:flex-row gap-4 mb-6 text-black">
                    <select
                        className="p-3 border border-gray-300 rounded-lg w-full focus:ring-2 focus:ring-indigo-400"
                        value={filtro.status}
                        onChange={(e) => setFiltro({ ...filtro, status: e.target.value })}
                    >
                        <option value="">Todos</option>
                        <option value="confirmado">Confirmado</option>
                        <option value="concluido">Concluído</option>
                        <option value="cancelado">Cancelado</option>
                    </select>

                    <select
                        className="p-3 border border-gray-300 rounded-lg w-full focus:ring-2 focus:ring-indigo-400"
                        value={filtro.ordenacao}
                        onChange={(e) => setFiltro({ ...filtro, ordenacao: e.target.value })}
                    >
                        <option value="recente">Mais recente</option>
                        <option value="distante">Mais distante</option>
                    </select>

                    <button
                        onClick={fetchData}
                        className="bg-laranja text-white font-bold p-3 rounded-lg w-full hover:bg-gray-400 transition"
                    >
                        Buscar
                    </button>
                </div>

                {/* Carregamento */}
                {loading && <p className="text-center text-gray-500">Carregando...</p>}

                {/* Lista de Agendamentos */}
                {agendamentos.length > 0 ? (
                    <div className="space-y-4">
                        {agendamentos.map((agendamento) => (
                            <div key={agendamento.id} className="bg-gray-100 p-5 rounded-xl shadow-md">
                                <h2 className="text-lg text-gray-900 font-archivo-black">{agendamento.servico.nome.toUpperCase()}</h2>
                                <p className="text-gray-700">{agendamento.servico.descricao}</p>
                                <p className="text-gray-600">
                                    Data:{" "}
                                    <span className="font-archivo-black">
                                        {new Date(agendamento.data_inicio).toLocaleString("pt-BR", {
                                            weekday: "long",
                                            day: "2-digit",
                                            month: "long",
                                            year: "numeric",
                                            hour: "2-digit",
                                            minute: "2-digit",
                                            second: "2-digit",
                                            hour12: false,
                                        }).toUpperCase()}
                                    </span>
                                </p>
                                <p className={`text-sm font-semibold ${
                                    agendamento.status === "confirmado" ? "text-green-600" :
                                    agendamento.status === "cancelado" ? "text-red-600" :
                                    "text-blue-600"
                                }`}>
                                    Status: {agendamento.status}
                                </p>
                            </div>
                        ))}
                    </div>
                ) : (
                    !loading && <p className="text-center text-gray-500">Nenhum agendamento encontrado.</p>
                )}
            </div>
        </section>
    );
}
