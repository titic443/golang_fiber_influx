package validators

import (
	"datalog-go/utils/logs"
	"encoding/json"
	"fmt"
	"strconv"
)

type Group2Dto struct {
	DATALOGID string  `json:"DATALOG ID" `
	BATTERYID string  `json:"BATTERY ID" `
	Soc       float64 `json:"Soc" `
	Cell1     float64 `json:"Cell_1" `
	Cell2     float64 `json:"Cell_2" `
	Cell3     float64 `json:"Cell_3" `
	Cell4     float64 `json:"Cell_4" `
	Cell5     float64 `json:"Cell_5" `
	Cell6     float64 `json:"Cell_6" `
	Cell7     float64 `json:"Cell_7" `
	Cell8     float64 `json:"Cell_8" `
	Cell9     float64 `json:"Cell_9" `
	Cell10    float64 `json:"Cell_10" `
	Cell11    float64 `json:"Cell_11" `
	Cell12    float64 `json:"Cell_12" `
	Cell13    float64 `json:"Cell_13" `
	Cell14    float64 `json:"Cell_14" `
	Cell15    float64 `json:"Cell_15" `
	Cell16    float64 `json:"Cell_16" `
	Cell17    float64 `json:"Cell_17" `
	Cell18    float64 `json:"Cell_18" `
	Cell19    float64 `json:"Cell_19" `
	Cell20    float64 `json:"Cell_20" `
	Cell21    float64 `json:"Cell_21" `
	Cell22    float64 `json:"Cell_22" `
	Cell23    float64 `json:"Cell_23" `
	Cell24    float64 `json:"Cell_24" `
	Cell25    float64 `json:"Cell_25" `
	Cell26    float64 `json:"Cell_26" `
	Cell27    float64 `json:"Cell_27" `
	Cell28    float64 `json:"Cell_28" `
	Cell29    float64 `json:"Cell_29" `
	Cell30    float64 `json:"Cell_30" `
	Cell31    float64 `json:"Cell_31" `
	Cell32    float64 `json:"Cell_32" `
	Cell33    float64 `json:"Cell_33" `
	Cell34    float64 `json:"Cell_34" `
	Cell35    float64 `json:"Cell_35" `
	Cell36    float64 `json:"Cell_36" `
	Cell37    float64 `json:"Cell_37" `
	Cell38    float64 `json:"Cell_38" `
	Cell39    float64 `json:"Cell_39" `
	Cell40    float64 `json:"Cell_40" `
	Cell41    float64 `json:"Cell_41" `
	Cell42    float64 `json:"Cell_42" `
	Cell43    float64 `json:"Cell_43" `
	Cell44    float64 `json:"Cell_44" `
	Cell45    float64 `json:"Cell_45" `
	Cell46    float64 `json:"Cell_46" `
	Cell47    float64 `json:"Cell_47" `
	Cell48    float64 `json:"Cell_48" `
	Cell49    float64 `json:"Cell_49" `
	Cell50    float64 `json:"Cell_50" `
	Cell51    float64 `json:"Cell_51" `
	Cell52    float64 `json:"Cell_52" `
	Cell53    float64 `json:"Cell_53" `
	Cell54    float64 `json:"Cell_54" `
	Cell55    float64 `json:"Cell_55" `
	Cell56    float64 `json:"Cell_56" `
	Cell57    float64 `json:"Cell_57" `
	Cell58    float64 `json:"Cell_58" `
	Cell59    float64 `json:"Cell_59" `
	Cell60    float64 `json:"Cell_60" `
	Cell61    float64 `json:"Cell_61" `
	Cell62    float64 `json:"Cell_62" `
	Cell63    float64 `json:"Cell_63" `
	Cell64    float64 `json:"Cell_64" `
	Cell65    float64 `json:"Cell_65" `
	Cell66    float64 `json:"Cell_66" `
	Cell67    float64 `json:"Cell_67" `
	Cell68    float64 `json:"Cell_68" `
	Cell69    float64 `json:"Cell_69" `
	Cell70    float64 `json:"Cell_70" `
	Cell71    float64 `json:"Cell_71" `
	Cell72    float64 `json:"Cell_72" `
	Cell73    float64 `json:"Cell_73" `
	Cell74    float64 `json:"Cell_74" `
	Cell75    float64 `json:"Cell_75" `
	Cell76    float64 `json:"Cell_76" `
	Cell77    float64 `json:"Cell_77" `
	Cell78    float64 `json:"Cell_78" `
	Cell79    float64 `json:"Cell_79" `
	Cell80    float64 `json:"Cell_80" `
	Cell81    float64 `json:"Cell_81" `
	Cell82    float64 `json:"Cell_82" `
	Cell83    float64 `json:"Cell_83" `
	Cell84    float64 `json:"Cell_84" `
	Cell85    float64 `json:"Cell_85" `
	Cell86    float64 `json:"Cell_86" `
	Cell87    float64 `json:"Cell_87" `
	Cell88    float64 `json:"Cell_88" `
	Cell89    float64 `json:"Cell_89" `
	Cell90    float64 `json:"Cell_90" `
	Cell91    float64 `json:"Cell_91" `
	Cell92    float64 `json:"Cell_92" `
	Cell93    float64 `json:"Cell_93" `
	Cell94    float64 `json:"Cell_94" `
	Cell95    float64 `json:"Cell_95" `
	Cell96    float64 `json:"Cell_96" `
	Cell97    float64 `json:"Cell_97" `
	Cell98    float64 `json:"Cell_98" `
	Cell99    float64 `json:"Cell_99" `
	Cell100   float64 `json:"Cell_100" `
	Cell101   float64 `json:"Cell_101" `
	Cell102   float64 `json:"Cell_102" `
	Cell103   float64 `json:"Cell_103" `
	Cell104   float64 `json:"Cell_104" `
	Cell105   float64 `json:"Cell_105" `
	Cell106   float64 `json:"Cell_106" `
	Cell107   float64 `json:"Cell_107" `
	Cell108   float64 `json:"Cell_108" `
	Cell109   float64 `json:"Cell_109" `
	Cell110   float64 `json:"Cell_110" `
	Cell111   float64 `json:"Cell_111" `
	Cell112   float64 `json:"Cell_112" `
	Cell113   float64 `json:"Cell_113" `
	Cell114   float64 `json:"Cell_114" `
	Cell115   float64 `json:"Cell_115" `
	Cell116   float64 `json:"Cell_116" `
	Cell117   float64 `json:"Cell_117" `
	Cell118   float64 `json:"Cell_118" `
	Cell119   float64 `json:"Cell_119" `
	Cell120   float64 `json:"Cell_120" `
	Cell121   float64 `json:"Cell_121" `
	Cell122   float64 `json:"Cell_122" `
	Cell123   float64 `json:"Cell_123" `
	Cell124   float64 `json:"Cell_124" `
	Cell125   float64 `json:"Cell_125" `
	Cell126   float64 `json:"Cell_126" `
	Cell127   float64 `json:"Cell_127" `
	Cell128   float64 `json:"Cell_128" `
	Cell129   float64 `json:"Cell_129" `
	Cell130   float64 `json:"Cell_130" `
	Cell131   float64 `json:"Cell_131" `
	Cell132   float64 `json:"Cell_132" `
	Cell133   float64 `json:"Cell_133" `
	Cell134   float64 `json:"Cell_134" `
	Cell135   float64 `json:"Cell_135" `
	Cell136   float64 `json:"Cell_136" `
	Cell137   float64 `json:"Cell_137" `
	Cell138   float64 `json:"Cell_138" `
	Cell139   float64 `json:"Cell_139" `
	Cell140   float64 `json:"Cell_140" `
	Cell141   float64 `json:"Cell_141" `
	Cell142   float64 `json:"Cell_142" `
	Cell143   float64 `json:"Cell_143" `
	Cell144   float64 `json:"Cell_144" `
	Cell145   float64 `json:"Cell_145" `
	Cell146   float64 `json:"Cell_146" `
	Cell147   float64 `json:"Cell_147" `
	Cell148   float64 `json:"Cell_148" `
	Cell149   float64 `json:"Cell_149" `
	Cell150   float64 `json:"Cell_150" `
	Cell151   float64 `json:"Cell_151" `
	Cell152   float64 `json:"Cell_152" `
	Cell153   float64 `json:"Cell_153" `
	Cell154   float64 `json:"Cell_154" `
	Cell155   float64 `json:"Cell_155" `
	Cell156   float64 `json:"Cell_156" `
	Cell157   float64 `json:"Cell_157" `
	Cell158   float64 `json:"Cell_158" `
	Cell159   float64 `json:"Cell_159" `
	Cell160   float64 `json:"Cell_160" `
	Cell161   float64 `json:"Cell_161" `
	Cell162   float64 `json:"Cell_162" `
	Cell163   float64 `json:"Cell_163" `
	Cell164   float64 `json:"Cell_164" `
	Cell165   float64 `json:"Cell_165" `
	Cell166   float64 `json:"Cell_166" `
	Cell167   float64 `json:"Cell_167" `
	Cell168   float64 `json:"Cell_168" `
	Cell169   float64 `json:"Cell_169" `
	Cell170   float64 `json:"Cell_170" `
	Cell171   float64 `json:"Cell_171" `
	Cell172   float64 `json:"Cell_172" `
	Cell173   float64 `json:"Cell_173" `
	Cell174   float64 `json:"Cell_174" `
	Cell175   float64 `json:"Cell_175" `
	Cell176   float64 `json:"Cell_176" `
	Cell177   float64 `json:"Cell_177" `
	Cell178   float64 `json:"Cell_178" `
	Cell179   float64 `json:"Cell_179" `
	Cell180   float64 `json:"Cell_180" `
	Cell181   float64 `json:"Cell_181" `
	Cell182   float64 `json:"Cell_182" `
	Cell183   float64 `json:"Cell_183" `
	Cell184   float64 `json:"Cell_184" `
	Cell185   float64 `json:"Cell_185" `
	Cell186   float64 `json:"Cell_186" `
	Cell187   float64 `json:"Cell_187" `
	Cell188   float64 `json:"Cell_188" `
	Cell189   float64 `json:"Cell_189" `
	Cell190   float64 `json:"Cell_190" `
	Cell191   float64 `json:"Cell_191" `
	Cell192   float64 `json:"Cell_192" `
	Cell193   float64 `json:"Cell_193" `
	Cell194   float64 `json:"Cell_194" `
	Cell195   float64 `json:"Cell_195" `
	Cell196   float64 `json:"Cell_196" `
	Cell197   float64 `json:"Cell_197" `
	Cell198   float64 `json:"Cell_198" `
	Cell199   float64 `json:"Cell_199" `
	Cell200   float64 `json:"Cell_200" `
	Cell201   float64 `json:"Cell_201" `
	Cell202   float64 `json:"Cell_202" `
	Cell203   float64 `json:"Cell_203" `
	Cell204   float64 `json:"Cell_204" `
	Cell205   float64 `json:"Cell_205" `
	Cell206   float64 `json:"Cell_206" `
	Cell207   float64 `json:"Cell_207" `
	Cell208   float64 `json:"Cell_208" `
	Cell209   float64 `json:"Cell_209" `
	Cell210   float64 `json:"Cell_210" `
	Cell211   float64 `json:"Cell_211" `
	Cell212   float64 `json:"Cell_212" `
	Cell213   float64 `json:"Cell_213" `
	Cell214   float64 `json:"Cell_214" `
	Cell215   float64 `json:"Cell_215" `
	Cell216   float64 `json:"Cell_216" `
	Cell217   float64 `json:"Cell_217" `
	Cell218   float64 `json:"Cell_218" `
	Cell219   float64 `json:"Cell_219" `
	Cell220   float64 `json:"Cell_220" `
	Cell221   float64 `json:"Cell_221" `
	Cell222   float64 `json:"Cell_222" `
	Cell223   float64 `json:"Cell_223" `
	Cell224   float64 `json:"Cell_224" `
	Cell225   float64 `json:"Cell_225" `
	Cell226   float64 `json:"Cell_226" `
	Cell227   float64 `json:"Cell_227" `
	Cell228   float64 `json:"Cell_228" `
	Cell229   float64 `json:"Cell_229" `
	Cell230   float64 `json:"Cell_230" `
	Cell231   float64 `json:"Cell_231" `
	Cell232   float64 `json:"Cell_232" `
	Cell233   float64 `json:"Cell_233" `
	Cell234   float64 `json:"Cell_234" `
	Cell235   float64 `json:"Cell_235" `
	Cell236   float64 `json:"Cell_236" `
	Cell237   float64 `json:"Cell_237" `
	Cell238   float64 `json:"Cell_238" `
	Cell239   float64 `json:"Cell_239" `
	Cell240   float64 `json:"Cell_240" `
	Cell241   float64 `json:"Cell_241" `
	Cell242   float64 `json:"Cell_242" `
	Cell243   float64 `json:"Cell_243" `
	Cell244   float64 `json:"Cell_244" `
	Cell245   float64 `json:"Cell_245" `
	Cell246   float64 `json:"Cell_246" `
	Cell247   float64 `json:"Cell_247" `
	Cell248   float64 `json:"Cell_248" `
	Cell249   float64 `json:"Cell_249" `
	Cell250   float64 `json:"Cell_250" `
	Cell251   float64 `json:"Cell_251" `
	Cell252   float64 `json:"Cell_252" `
	Cell253   float64 `json:"Cell_253" `
	Cell254   float64 `json:"Cell_254" `
	Cell255   float64 `json:"Cell_255" `
	Cell256   float64 `json:"Cell_256" `
	Cell257   float64 `json:"Cell_257" `
	Cell258   float64 `json:"Cell_258" `
	Cell259   float64 `json:"Cell_259" `
	Cell260   float64 `json:"Cell_260" `
	Cell261   float64 `json:"Cell_261" `
	Cell262   float64 `json:"Cell_262" `
	Cell263   float64 `json:"Cell_263" `
	Cell264   float64 `json:"Cell_264" `
	Cell265   float64 `json:"Cell_265" `
	Cell266   float64 `json:"Cell_266" `
	Cell267   float64 `json:"Cell_267" `
	Cell268   float64 `json:"Cell_268" `
	Cell269   float64 `json:"Cell_269" `
	Cell270   float64 `json:"Cell_270" `
	Cell271   float64 `json:"Cell_271" `
	Cell272   float64 `json:"Cell_272" `
	Cell273   float64 `json:"Cell_273" `
	Cell274   float64 `json:"Cell_274" `
	Cell275   float64 `json:"Cell_275" `
	Cell276   float64 `json:"Cell_276" `
	Cell277   float64 `json:"Cell_277" `
	Cell278   float64 `json:"Cell_278" `
	Cell279   float64 `json:"Cell_279" `
	Cell280   float64 `json:"Cell_280" `
	Cell281   float64 `json:"Cell_281" `
	Cell282   float64 `json:"Cell_282" `
	Cell283   float64 `json:"Cell_283" `
	Cell284   float64 `json:"Cell_284" `
	Cell285   float64 `json:"Cell_285" `
	Cell286   float64 `json:"Cell_286" `
	Cell287   float64 `json:"Cell_287" `
	Cell288   float64 `json:"Cell_288" `
	Cell289   float64 `json:"Cell_289" `
	Cell290   float64 `json:"Cell_290" `
	Cell291   float64 `json:"Cell_291" `
	Cell292   float64 `json:"Cell_292" `
	Cell293   float64 `json:"Cell_293" `
	Cell294   float64 `json:"Cell_294" `
	Cell295   float64 `json:"Cell_295" `
	Cell296   float64 `json:"Cell_296" `
	Cell297   float64 `json:"Cell_297" `
	Cell298   float64 `json:"Cell_298" `
	Cell299   float64 `json:"Cell_299" `
	Cell300   float64 `json:"Cell_300" `
	Cell301   float64 `json:"Cell_301" `
	Cell302   float64 `json:"Cell_302" `
	Cell303   float64 `json:"Cell_303" `
	Cell304   float64 `json:"Cell_304" `
	Cell305   float64 `json:"Cell_305" `
	Cell306   float64 `json:"Cell_306" `
	Cell307   float64 `json:"Cell_307" `
	Cell308   float64 `json:"Cell_308" `
	Cell309   float64 `json:"Cell_309" `
	Cell310   float64 `json:"Cell_310" `
	Cell311   float64 `json:"Cell_311" `
	Cell312   float64 `json:"Cell_312" `
	Cell313   float64 `json:"Cell_313" `
	Cell314   float64 `json:"Cell_314" `
	Cell315   float64 `json:"Cell_315" `
	Cell316   float64 `json:"Cell_316" `
	Cell317   float64 `json:"Cell_317" `
	Cell318   float64 `json:"Cell_318" `
	Cell319   float64 `json:"Cell_319" `
	Cell320   float64 `json:"Cell_320" `
	Cell321   float64 `json:"Cell_321" `
	Cell322   float64 `json:"Cell_322" `
	Cell323   float64 `json:"Cell_323" `
	Cell324   float64 `json:"Cell_324" `
	Cell325   float64 `json:"Cell_325" `
	Cell326   float64 `json:"Cell_326" `
	Cell327   float64 `json:"Cell_327" `
	Cell328   float64 `json:"Cell_328" `
	Cell329   float64 `json:"Cell_329" `
	Cell330   float64 `json:"Cell_330" `
	Cell331   float64 `json:"Cell_331" `
	Cell332   float64 `json:"Cell_332" `
	Cell333   float64 `json:"Cell_333" `
	Cell334   float64 `json:"Cell_334" `
	Cell335   float64 `json:"Cell_335" `
	Cell336   float64 `json:"Cell_336" `
	BlQM1     float64 `json:"BL_Q_M1" `
	BlQM2     float64 `json:"BL_Q_M2" `
	BlQM3     float64 `json:"BL_Q_M3" `
	BlQM4     float64 `json:"BL_Q_M4" `
	BlQM5     float64 `json:"BL_Q_M5" `
	BlQM6     float64 `json:"BL_Q_M6" `
	BlQM7     float64 `json:"BL_Q_M7" `
	BlQM8     float64 `json:"BL_Q_M8" `
	BlQM9     float64 `json:"BL_Q_M9" `
	BlQM10    float64 `json:"BL_Q_M10" `
	BlQM11    float64 `json:"BL_Q_M11" `
	BlQM12    float64 `json:"BL_Q_M12" `
	BlQM13    float64 `json:"BL_Q_M13" `
	BlQM14    float64 `json:"BL_Q_M14" `
	BlQM15    float64 `json:"BL_Q_M15" `
	BlQM16    float64 `json:"BL_Q_M16" `
	BlQM17    float64 `json:"BL_Q_M17" `
	BlQM18    float64 `json:"BL_Q_M18" `
	BlQM19    float64 `json:"BL_Q_M19" `
	BlQM20    float64 `json:"BL_Q_M20" `
	BlQM21    float64 `json:"BL_Q_M21" `
	BlQM22    float64 `json:"BL_Q_M22" `
	BlQM23    float64 `json:"BL_Q_M23" `
	BlQM24    float64 `json:"BL_Q_M24" `
	NewBLM1   float64 `json:"New_BL_M1" `
	NewBLM2   float64 `json:"New_BL_M2" `
	NewBLM3   float64 `json:"New_BL_M3" `
	NewBLM4   float64 `json:"New_BL_M4" `
	NewBLM5   float64 `json:"New_BL_M5" `
	NewBLM6   float64 `json:"New_BL_M6" `
	NewBLM7   float64 `json:"New_BL_M7" `
	NewBLM8   float64 `json:"New_BL_M8" `
	NewBLM9   float64 `json:"New_BL_M9" `
	NewBLM10  float64 `json:"New_BL_M10" `
	NewBLM11  float64 `json:"New_BL_M11" `
	NewBLM12  float64 `json:"New_BL_M12" `
	NewBLM13  float64 `json:"New_BL_M13" `
	NewBLM14  float64 `json:"New_BL_M14" `
	NewBLM15  float64 `json:"New_BL_M15" `
	NewBLM16  float64 `json:"New_BL_M16" `
	NewBLM17  float64 `json:"New_BL_M17" `
	NewBLM18  float64 `json:"New_BL_M18" `
	NewBLM19  float64 `json:"New_BL_M19" `
	NewBLM20  float64 `json:"New_BL_M20" `
	NewBLM21  float64 `json:"New_BL_M21" `
	NewBLM22  float64 `json:"New_BL_M22" `
	NewBLM23  float64 `json:"New_BL_M23" `
	NewBLM24  float64 `json:"New_BL_M24" `
}

func (g *Group2Dto) MapType(b []byte) []interface{} {
	var con []interface{}
	var m []map[string]string
	metric := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		logs.Error(err)
	}
	for _, v := range m {
		t := Group2Dto{
			DATALOGID: v["DATALOG ID"],
			BATTERYID: v["BATTERY ID"],
		}

		for k, vv := range v {
			if k != "DATALOG ID" && k != "BATTERY ID" {

				f, err := strconv.ParseFloat(vv, 64)
				if err != nil {
					metric[k] = 0
				}
				metric[k] = f
			}
		}

		b, err := json.Marshal(metric)
		if err != nil {
			logs.Error(err)
		}
		json.Unmarshal(b, &t)
		con = append(con, t)
		fmt.Println(con...)
	}

	return con
}
