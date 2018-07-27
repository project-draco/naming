package naming

import (
	"testing"
)

func TestJavaToHR(t *testing.T) {
	for _, data := range []struct{ input, output string }{
		{
			"br.mil.eb.cds.sisbol.boletim.fabrica.ResourceFactory.getEntityManager()",
			"br_mil_eb_cds_sisbol_boletim_fabrica_ResourceFactory.java/[CN]/ResourceFactory/[MT]/getEntityManager()"},
		{
			"br.mil.eb.cds.sisbol.boletim.modelo.Alteracao.setAno(java.lang.Integer)",
			"br_mil_eb_cds_sisbol_boletim_modelo_Alteracao.java/[CN]/Alteracao/[MT]/setAno(Integer)",
		},
		{
			"br.mil.eb.cds.sisbol.boletim.repositorio.AlteracaoRepositorioJPA.buscarTSPorIdentidadeAnoSemestre(java.lang.String, int, int)",
			"br_mil_eb_cds_sisbol_boletim_repositorio_AlteracaoRepositorioJPA.java/[CN]/AlteracaoRepositorioJPA/[MT]/buscarTSPorIdentidadeAnoSemestre(String,int,int)",
		},
		{
			"br.mil.eb.cds.sisbol.boletim.modelo.AssuntoGeral.secaoBoletim",
			"br_mil_eb_cds_sisbol_boletim_modelo_AssuntoGeral.java/[CN]/AssuntoGeral/[FE]/secaoBoletim",
		},
		{
			"br.mil.eb.cds.sisbol.boletim.modelo.Boletim.Boletim()",
			"br_mil_eb_cds_sisbol_boletim_modelo_Boletim.java/[CN]/Boletim/[CS]/Boletim()",
		},
		{
			"br.mil.eb.cds.sisbol.boletim.modelo.BoletimStatusEnum.static {}",
			"",
		},
	} {
		if result := JavaToHR(data.input); result != data.output {
			t.Errorf("Got %v want %v", result, data.output)
		}
	}
}
