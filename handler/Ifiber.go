package handlers

type Group1Dto struct {
	DATALOGID string  `json:"DATALOG ID" validate:"required"`
	BATTERYID string  `json:"BATTERY ID" validate:"required"`
	TotalVV   float64 `json:"Total V (V),string" validate:"required"`
	TotalAA   float64 `json:"Total A (A),string" validate:"required"`
	MinVMV    float64 `json:"MinV (mV),string" validate:"required"`
	MaxVMV    float64 `json:"MaxV (mV),string" validate:"required"`
}

type Group2Dto struct {
	DATALOGID string  `json:"DATALOG ID,string" validate:"required"`
	BATTERYID string  `json:"BATTERY ID,string" validate:"required"`
	Soc       float64 `json:"Soc,string" validate:"required"`
	Cell1     float64 `json:"Cell_1,string" validate:"required"`
	Cell2     float64 `json:"Cell_2,string" validate:"required"`
	Cell3     float64 `json:"Cell_3,string" validate:"required"`
	Cell4     float64 `json:"Cell_4,string" validate:"required"`
	Cell5     float64 `json:"Cell_5,string" validate:"required"`
	Cell6     float64 `json:"Cell_6,string" validate:"required"`
	Cell7     float64 `json:"Cell_7,string" validate:"required"`
	Cell8     float64 `json:"Cell_8,string" validate:"required"`
	Cell9     float64 `json:"Cell_9,string" validate:"required"`
	Cell10    float64 `json:"Cell_10,string" validate:"required"`
	Cell11    float64 `json:"Cell_11,string" validate:"required"`
	Cell12    float64 `json:"Cell_12,string" validate:"required"`
	Cell13    float64 `json:"Cell_13,string" validate:"required"`
	Cell14    float64 `json:"Cell_14,string" validate:"required"`
	Cell15    float64 `json:"Cell_15,string" validate:"required"`
	Cell16    float64 `json:"Cell_16,string" validate:"required"`
	Cell17    float64 `json:"Cell_17,string" validate:"required"`
	Cell18    float64 `json:"Cell_18,string" validate:"required"`
	Cell19    float64 `json:"Cell_19,string" validate:"required"`
	Cell20    float64 `json:"Cell_20,string" validate:"required"`
	Cell21    float64 `json:"Cell_21,string" validate:"required"`
	Cell22    float64 `json:"Cell_22,string" validate:"required"`
	Cell23    float64 `json:"Cell_23,string" validate:"required"`
	Cell24    float64 `json:"Cell_24,string" validate:"required"`
	Cell25    float64 `json:"Cell_25,string" validate:"required"`
	Cell26    float64 `json:"Cell_26,string" validate:"required"`
	Cell27    float64 `json:"Cell_27,string" validate:"required"`
	Cell28    float64 `json:"Cell_28,string" validate:"required"`
	Cell29    float64 `json:"Cell_29,string" validate:"required"`
	Cell30    float64 `json:"Cell_30,string" validate:"required"`
	Cell31    float64 `json:"Cell_31,string" validate:"required"`
	Cell32    float64 `json:"Cell_32,string" validate:"required"`
	Cell33    float64 `json:"Cell_33,string" validate:"required"`
	Cell34    float64 `json:"Cell_34,string" validate:"required"`
	Cell35    float64 `json:"Cell_35,string" validate:"required"`
	Cell36    float64 `json:"Cell_36,string" validate:"required"`
	Cell37    float64 `json:"Cell_37,string" validate:"required"`
	Cell38    float64 `json:"Cell_38,string" validate:"required"`
	Cell39    float64 `json:"Cell_39,string" validate:"required"`
	Cell40    float64 `json:"Cell_40,string" validate:"required"`
	Cell41    float64 `json:"Cell_41,string" validate:"required"`
	Cell42    float64 `json:"Cell_42,string" validate:"required"`
	Cell43    float64 `json:"Cell_43,string" validate:"required"`
	Cell44    float64 `json:"Cell_44,string" validate:"required"`
	Cell45    float64 `json:"Cell_45,string" validate:"required"`
	Cell46    float64 `json:"Cell_46,string" validate:"required"`
	Cell47    float64 `json:"Cell_47,string" validate:"required"`
	Cell48    float64 `json:"Cell_48,string" validate:"required"`
	Cell49    float64 `json:"Cell_49,string" validate:"required"`
	Cell50    float64 `json:"Cell_50,string" validate:"required"`
	Cell51    float64 `json:"Cell_51,string" validate:"required"`
	Cell52    float64 `json:"Cell_52,string" validate:"required"`
	Cell53    float64 `json:"Cell_53,string" validate:"required"`
	Cell54    float64 `json:"Cell_54,string" validate:"required"`
	Cell55    float64 `json:"Cell_55,string" validate:"required"`
	Cell56    float64 `json:"Cell_56,string" validate:"required"`
	Cell57    float64 `json:"Cell_57,string" validate:"required"`
	Cell58    float64 `json:"Cell_58,string" validate:"required"`
	Cell59    float64 `json:"Cell_59,string" validate:"required"`
	Cell60    float64 `json:"Cell_60,string" validate:"required"`
	Cell61    float64 `json:"Cell_61,string" validate:"required"`
	Cell62    float64 `json:"Cell_62,string" validate:"required"`
	Cell63    float64 `json:"Cell_63,string" validate:"required"`
	Cell64    float64 `json:"Cell_64,string" validate:"required"`
	Cell65    float64 `json:"Cell_65,string" validate:"required"`
	Cell66    float64 `json:"Cell_66,string" validate:"required"`
	Cell67    float64 `json:"Cell_67,string" validate:"required"`
	Cell68    float64 `json:"Cell_68,string" validate:"required"`
	Cell69    float64 `json:"Cell_69,string" validate:"required"`
	Cell70    float64 `json:"Cell_70,string" validate:"required"`
	Cell71    float64 `json:"Cell_71,string" validate:"required"`
	Cell72    float64 `json:"Cell_72,string" validate:"required"`
	Cell73    float64 `json:"Cell_73,string" validate:"required"`
	Cell74    float64 `json:"Cell_74,string" validate:"required"`
	Cell75    float64 `json:"Cell_75,string" validate:"required"`
	Cell76    float64 `json:"Cell_76,string" validate:"required"`
	Cell77    float64 `json:"Cell_77,string" validate:"required"`
	Cell78    float64 `json:"Cell_78,string" validate:"required"`
	Cell79    float64 `json:"Cell_79,string" validate:"required"`
	Cell80    float64 `json:"Cell_80,string" validate:"required"`
	Cell81    float64 `json:"Cell_81,string" validate:"required"`
	Cell82    float64 `json:"Cell_82,string" validate:"required"`
	Cell83    float64 `json:"Cell_83,string" validate:"required"`
	Cell84    float64 `json:"Cell_84,string" validate:"required"`
	Cell85    float64 `json:"Cell_85,string" validate:"required"`
	Cell86    float64 `json:"Cell_86,string" validate:"required"`
	Cell87    float64 `json:"Cell_87,string" validate:"required"`
	Cell88    float64 `json:"Cell_88,string" validate:"required"`
	Cell89    float64 `json:"Cell_89,string" validate:"required"`
	Cell90    float64 `json:"Cell_90,string" validate:"required"`
	Cell91    float64 `json:"Cell_91,string" validate:"required"`
	Cell92    float64 `json:"Cell_92,string" validate:"required"`
	Cell93    float64 `json:"Cell_93,string" validate:"required"`
	Cell94    float64 `json:"Cell_94,string" validate:"required"`
	Cell95    float64 `json:"Cell_95,string" validate:"required"`
	Cell96    float64 `json:"Cell_96,string" validate:"required"`
	Cell97    float64 `json:"Cell_97,string" validate:"required"`
	Cell98    float64 `json:"Cell_98,string" validate:"required"`
	Cell99    float64 `json:"Cell_99,string" validate:"required"`
	Cell100   float64 `json:"Cell_100,string" validate:"required"`
	Cell101   float64 `json:"Cell_101,string" validate:"required"`
	Cell102   float64 `json:"Cell_102,string" validate:"required"`
	Cell103   float64 `json:"Cell_103,string" validate:"required"`
	Cell104   float64 `json:"Cell_104,string" validate:"required"`
	Cell105   float64 `json:"Cell_105,string" validate:"required"`
	Cell106   float64 `json:"Cell_106,string" validate:"required"`
	Cell107   float64 `json:"Cell_107,string" validate:"required"`
	Cell108   float64 `json:"Cell_108,string" validate:"required"`
	Cell109   float64 `json:"Cell_109,string" validate:"required"`
	Cell110   float64 `json:"Cell_110,string" validate:"required"`
	Cell111   float64 `json:"Cell_111,string" validate:"required"`
	Cell112   float64 `json:"Cell_112,string" validate:"required"`
	Cell113   float64 `json:"Cell_113,string" validate:"required"`
	Cell114   float64 `json:"Cell_114,string" validate:"required"`
	Cell115   float64 `json:"Cell_115,string" validate:"required"`
	Cell116   float64 `json:"Cell_116,string" validate:"required"`
	Cell117   float64 `json:"Cell_117,string" validate:"required"`
	Cell118   float64 `json:"Cell_118,string" validate:"required"`
	Cell119   float64 `json:"Cell_119,string" validate:"required"`
	Cell120   float64 `json:"Cell_120,string" validate:"required"`
	Cell121   float64 `json:"Cell_121,string" validate:"required"`
	Cell122   float64 `json:"Cell_122,string" validate:"required"`
	Cell123   float64 `json:"Cell_123,string" validate:"required"`
	Cell124   float64 `json:"Cell_124,string" validate:"required"`
	Cell125   float64 `json:"Cell_125,string" validate:"required"`
	Cell126   float64 `json:"Cell_126,string" validate:"required"`
	Cell127   float64 `json:"Cell_127,string" validate:"required"`
	Cell128   float64 `json:"Cell_128,string" validate:"required"`
	Cell129   float64 `json:"Cell_129,string" validate:"required"`
	Cell130   float64 `json:"Cell_130,string" validate:"required"`
	Cell131   float64 `json:"Cell_131,string" validate:"required"`
	Cell132   float64 `json:"Cell_132,string" validate:"required"`
	Cell133   float64 `json:"Cell_133,string" validate:"required"`
	Cell134   float64 `json:"Cell_134,string" validate:"required"`
	Cell135   float64 `json:"Cell_135,string" validate:"required"`
	Cell136   float64 `json:"Cell_136,string" validate:"required"`
	Cell137   float64 `json:"Cell_137,string" validate:"required"`
	Cell138   float64 `json:"Cell_138,string" validate:"required"`
	Cell139   float64 `json:"Cell_139,string" validate:"required"`
	Cell140   float64 `json:"Cell_140,string" validate:"required"`
	Cell141   float64 `json:"Cell_141,string" validate:"required"`
	Cell142   float64 `json:"Cell_142,string" validate:"required"`
	Cell143   float64 `json:"Cell_143,string" validate:"required"`
	Cell144   float64 `json:"Cell_144,string" validate:"required"`
	Cell145   float64 `json:"Cell_145,string" validate:"required"`
	Cell146   float64 `json:"Cell_146,string" validate:"required"`
	Cell147   float64 `json:"Cell_147,string" validate:"required"`
	Cell148   float64 `json:"Cell_148,string" validate:"required"`
	Cell149   float64 `json:"Cell_149,string" validate:"required"`
	Cell150   float64 `json:"Cell_150,string" validate:"required"`
	Cell151   float64 `json:"Cell_151,string" validate:"required"`
	Cell152   float64 `json:"Cell_152,string" validate:"required"`
	Cell153   float64 `json:"Cell_153,string" validate:"required"`
	Cell154   float64 `json:"Cell_154,string" validate:"required"`
	Cell155   float64 `json:"Cell_155,string" validate:"required"`
	Cell156   float64 `json:"Cell_156,string" validate:"required"`
	Cell157   float64 `json:"Cell_157,string" validate:"required"`
	Cell158   float64 `json:"Cell_158,string" validate:"required"`
	Cell159   float64 `json:"Cell_159,string" validate:"required"`
	Cell160   float64 `json:"Cell_160,string" validate:"required"`
	Cell161   float64 `json:"Cell_161,string" validate:"required"`
	Cell162   float64 `json:"Cell_162,string" validate:"required"`
	Cell163   float64 `json:"Cell_163,string" validate:"required"`
	Cell164   float64 `json:"Cell_164,string" validate:"required"`
	Cell165   float64 `json:"Cell_165,string" validate:"required"`
	Cell166   float64 `json:"Cell_166,string" validate:"required"`
	Cell167   float64 `json:"Cell_167,string" validate:"required"`
	Cell168   float64 `json:"Cell_168,string" validate:"required"`
	Cell169   float64 `json:"Cell_169,string" validate:"required"`
	Cell170   float64 `json:"Cell_170,string" validate:"required"`
	Cell171   float64 `json:"Cell_171,string" validate:"required"`
	Cell172   float64 `json:"Cell_172,string" validate:"required"`
	Cell173   float64 `json:"Cell_173,string" validate:"required"`
	Cell174   float64 `json:"Cell_174,string" validate:"required"`
	Cell175   float64 `json:"Cell_175,string" validate:"required"`
	Cell176   float64 `json:"Cell_176,string" validate:"required"`
	Cell177   float64 `json:"Cell_177,string" validate:"required"`
	Cell178   float64 `json:"Cell_178,string" validate:"required"`
	Cell179   float64 `json:"Cell_179,string" validate:"required"`
	Cell180   float64 `json:"Cell_180,string" validate:"required"`
	Cell181   float64 `json:"Cell_181,string" validate:"required"`
	Cell182   float64 `json:"Cell_182,string" validate:"required"`
	Cell183   float64 `json:"Cell_183,string" validate:"required"`
	Cell184   float64 `json:"Cell_184,string" validate:"required"`
	Cell185   float64 `json:"Cell_185,string" validate:"required"`
	Cell186   float64 `json:"Cell_186,string" validate:"required"`
	Cell187   float64 `json:"Cell_187,string" validate:"required"`
	Cell188   float64 `json:"Cell_188,string" validate:"required"`
	Cell189   float64 `json:"Cell_189,string" validate:"required"`
	Cell190   float64 `json:"Cell_190,string" validate:"required"`
	Cell191   float64 `json:"Cell_191,string" validate:"required"`
	Cell192   float64 `json:"Cell_192,string" validate:"required"`
	Cell193   float64 `json:"Cell_193,string" validate:"required"`
	Cell194   float64 `json:"Cell_194,string" validate:"required"`
	Cell195   float64 `json:"Cell_195,string" validate:"required"`
	Cell196   float64 `json:"Cell_196,string" validate:"required"`
	Cell197   float64 `json:"Cell_197,string" validate:"required"`
	Cell198   float64 `json:"Cell_198,string" validate:"required"`
	Cell199   float64 `json:"Cell_199,string" validate:"required"`
	Cell200   float64 `json:"Cell_200,string" validate:"required"`
	Cell201   float64 `json:"Cell_201,string" validate:"required"`
	Cell202   float64 `json:"Cell_202,string" validate:"required"`
	Cell203   float64 `json:"Cell_203,string" validate:"required"`
	Cell204   float64 `json:"Cell_204,string" validate:"required"`
	Cell205   float64 `json:"Cell_205,string" validate:"required"`
	Cell206   float64 `json:"Cell_206,string" validate:"required"`
	Cell207   float64 `json:"Cell_207,string" validate:"required"`
	Cell208   float64 `json:"Cell_208,string" validate:"required"`
	Cell209   float64 `json:"Cell_209,string" validate:"required"`
	Cell210   float64 `json:"Cell_210,string" validate:"required"`
	Cell211   float64 `json:"Cell_211,string" validate:"required"`
	Cell212   float64 `json:"Cell_212,string" validate:"required"`
	Cell213   float64 `json:"Cell_213,string" validate:"required"`
	Cell214   float64 `json:"Cell_214,string" validate:"required"`
	Cell215   float64 `json:"Cell_215,string" validate:"required"`
	Cell216   float64 `json:"Cell_216,string" validate:"required"`
	Cell217   float64 `json:"Cell_217,string" validate:"required"`
	Cell218   float64 `json:"Cell_218,string" validate:"required"`
	Cell219   float64 `json:"Cell_219,string" validate:"required"`
	Cell220   float64 `json:"Cell_220,string" validate:"required"`
	Cell221   float64 `json:"Cell_221,string" validate:"required"`
	Cell222   float64 `json:"Cell_222,string" validate:"required"`
	Cell223   float64 `json:"Cell_223,string" validate:"required"`
	Cell224   float64 `json:"Cell_224,string" validate:"required"`
	Cell225   float64 `json:"Cell_225,string" validate:"required"`
	Cell226   float64 `json:"Cell_226,string" validate:"required"`
	Cell227   float64 `json:"Cell_227,string" validate:"required"`
	Cell228   float64 `json:"Cell_228,string" validate:"required"`
	Cell229   float64 `json:"Cell_229,string" validate:"required"`
	Cell230   float64 `json:"Cell_230,string" validate:"required"`
	Cell231   float64 `json:"Cell_231,string" validate:"required"`
	Cell232   float64 `json:"Cell_232,string" validate:"required"`
	Cell233   float64 `json:"Cell_233,string" validate:"required"`
	Cell234   float64 `json:"Cell_234,string" validate:"required"`
	Cell235   float64 `json:"Cell_235,string" validate:"required"`
	Cell236   float64 `json:"Cell_236,string" validate:"required"`
	Cell237   float64 `json:"Cell_237,string" validate:"required"`
	Cell238   float64 `json:"Cell_238,string" validate:"required"`
	Cell239   float64 `json:"Cell_239,string" validate:"required"`
	Cell240   float64 `json:"Cell_240,string" validate:"required"`
	Cell241   float64 `json:"Cell_241,string" validate:"required"`
	Cell242   float64 `json:"Cell_242,string" validate:"required"`
	Cell243   float64 `json:"Cell_243,string" validate:"required"`
	Cell244   float64 `json:"Cell_244,string" validate:"required"`
	Cell245   float64 `json:"Cell_245,string" validate:"required"`
	Cell246   float64 `json:"Cell_246,string" validate:"required"`
	Cell247   float64 `json:"Cell_247,string" validate:"required"`
	Cell248   float64 `json:"Cell_248,string" validate:"required"`
	Cell249   float64 `json:"Cell_249,string" validate:"required"`
	Cell250   float64 `json:"Cell_250,string" validate:"required"`
	Cell251   float64 `json:"Cell_251,string" validate:"required"`
	Cell252   float64 `json:"Cell_252,string" validate:"required"`
	Cell253   float64 `json:"Cell_253,string" validate:"required"`
	Cell254   float64 `json:"Cell_254,string" validate:"required"`
	Cell255   float64 `json:"Cell_255,string" validate:"required"`
	Cell256   float64 `json:"Cell_256,string" validate:"required"`
	Cell257   float64 `json:"Cell_257,string" validate:"required"`
	Cell258   float64 `json:"Cell_258,string" validate:"required"`
	Cell259   float64 `json:"Cell_259,string" validate:"required"`
	Cell260   float64 `json:"Cell_260,string" validate:"required"`
	Cell261   float64 `json:"Cell_261,string" validate:"required"`
	Cell262   float64 `json:"Cell_262,string" validate:"required"`
	Cell263   float64 `json:"Cell_263,string" validate:"required"`
	Cell264   float64 `json:"Cell_264,string" validate:"required"`
	Cell265   float64 `json:"Cell_265,string" validate:"required"`
	Cell266   float64 `json:"Cell_266,string" validate:"required"`
	Cell267   float64 `json:"Cell_267,string" validate:"required"`
	Cell268   float64 `json:"Cell_268,string" validate:"required"`
	Cell269   float64 `json:"Cell_269,string" validate:"required"`
	Cell270   float64 `json:"Cell_270,string" validate:"required"`
	Cell271   float64 `json:"Cell_271,string" validate:"required"`
	Cell272   float64 `json:"Cell_272,string" validate:"required"`
	Cell273   float64 `json:"Cell_273,string" validate:"required"`
	Cell274   float64 `json:"Cell_274,string" validate:"required"`
	Cell275   float64 `json:"Cell_275,string" validate:"required"`
	Cell276   float64 `json:"Cell_276,string" validate:"required"`
	Cell277   float64 `json:"Cell_277,string" validate:"required"`
	Cell278   float64 `json:"Cell_278,string" validate:"required"`
	Cell279   float64 `json:"Cell_279,string" validate:"required"`
	Cell280   float64 `json:"Cell_280,string" validate:"required"`
	Cell281   float64 `json:"Cell_281,string" validate:"required"`
	Cell282   float64 `json:"Cell_282,string" validate:"required"`
	Cell283   float64 `json:"Cell_283,string" validate:"required"`
	Cell284   float64 `json:"Cell_284,string" validate:"required"`
	Cell285   float64 `json:"Cell_285,string" validate:"required"`
	Cell286   float64 `json:"Cell_286,string" validate:"required"`
	Cell287   float64 `json:"Cell_287,string" validate:"required"`
	Cell288   float64 `json:"Cell_288,string" validate:"required"`
	Cell289   float64 `json:"Cell_289,string" validate:"required"`
	Cell290   float64 `json:"Cell_290,string" validate:"required"`
	Cell291   float64 `json:"Cell_291,string" validate:"required"`
	Cell292   float64 `json:"Cell_292,string" validate:"required"`
	Cell293   float64 `json:"Cell_293,string" validate:"required"`
	Cell294   float64 `json:"Cell_294,string" validate:"required"`
	Cell295   float64 `json:"Cell_295,string" validate:"required"`
	Cell296   float64 `json:"Cell_296,string" validate:"required"`
	Cell297   float64 `json:"Cell_297,string" validate:"required"`
	Cell298   float64 `json:"Cell_298,string" validate:"required"`
	Cell299   float64 `json:"Cell_299,string" validate:"required"`
	Cell300   float64 `json:"Cell_300,string" validate:"required"`
	Cell301   float64 `json:"Cell_301,string" validate:"required"`
	Cell302   float64 `json:"Cell_302,string" validate:"required"`
	Cell303   float64 `json:"Cell_303,string" validate:"required"`
	Cell304   float64 `json:"Cell_304,string" validate:"required"`
	Cell305   float64 `json:"Cell_305,string" validate:"required"`
	Cell306   float64 `json:"Cell_306,string" validate:"required"`
	Cell307   float64 `json:"Cell_307,string" validate:"required"`
	Cell308   float64 `json:"Cell_308,string" validate:"required"`
	Cell309   float64 `json:"Cell_309,string" validate:"required"`
	Cell310   float64 `json:"Cell_310,string" validate:"required"`
	Cell311   float64 `json:"Cell_311,string" validate:"required"`
	Cell312   float64 `json:"Cell_312,string" validate:"required"`
	Cell313   float64 `json:"Cell_313,string" validate:"required"`
	Cell314   float64 `json:"Cell_314,string" validate:"required"`
	Cell315   float64 `json:"Cell_315,string" validate:"required"`
	Cell316   float64 `json:"Cell_316,string" validate:"required"`
	Cell317   float64 `json:"Cell_317,string" validate:"required"`
	Cell318   float64 `json:"Cell_318,string" validate:"required"`
	Cell319   float64 `json:"Cell_319,string" validate:"required"`
	Cell320   float64 `json:"Cell_320,string" validate:"required"`
	Cell321   float64 `json:"Cell_321,string" validate:"required"`
	Cell322   float64 `json:"Cell_322,string" validate:"required"`
	Cell323   float64 `json:"Cell_323,string" validate:"required"`
	Cell324   float64 `json:"Cell_324,string" validate:"required"`
	Cell325   float64 `json:"Cell_325,string" validate:"required"`
	Cell326   float64 `json:"Cell_326,string" validate:"required"`
	Cell327   float64 `json:"Cell_327,string" validate:"required"`
	Cell328   float64 `json:"Cell_328,string" validate:"required"`
	Cell329   float64 `json:"Cell_329,string" validate:"required"`
	Cell330   float64 `json:"Cell_330,string" validate:"required"`
	Cell331   float64 `json:"Cell_331,string" validate:"required"`
	Cell332   float64 `json:"Cell_332,string" validate:"required"`
	Cell333   float64 `json:"Cell_333,string" validate:"required"`
	Cell334   float64 `json:"Cell_334,string" validate:"required"`
	Cell335   float64 `json:"Cell_335,string" validate:"required"`
	Cell336   float64 `json:"Cell_336,string" validate:"required"`
	BlQM1     float64 `json:"BL_Q_M1,string" validate:"required"`
	BlQM2     float64 `json:"BL_Q_M2,string" validate:"required"`
	BlQM3     float64 `json:"BL_Q_M3,string" validate:"required"`
	BlQM4     float64 `json:"BL_Q_M4,string" validate:"required"`
	BlQM5     float64 `json:"BL_Q_M5,string" validate:"required"`
	BlQM6     float64 `json:"BL_Q_M6,string" validate:"required"`
	BlQM7     float64 `json:"BL_Q_M7,string" validate:"required"`
	BlQM8     float64 `json:"BL_Q_M8,string" validate:"required"`
	BlQM9     float64 `json:"BL_Q_M9,string" validate:"required"`
	BlQM10    float64 `json:"BL_Q_M10,string" validate:"required"`
	BlQM11    float64 `json:"BL_Q_M11,string" validate:"required"`
	BlQM12    float64 `json:"BL_Q_M12,string" validate:"required"`
	BlQM13    float64 `json:"BL_Q_M13,string" validate:"required"`
	BlQM14    float64 `json:"BL_Q_M14,string" validate:"required"`
	BlQM15    float64 `json:"BL_Q_M15,string" validate:"required"`
	BlQM16    float64 `json:"BL_Q_M16,string" validate:"required"`
	BlQM17    float64 `json:"BL_Q_M17,string" validate:"required"`
	BlQM18    float64 `json:"BL_Q_M18,string" validate:"required"`
	BlQM19    float64 `json:"BL_Q_M19,string" validate:"required"`
	BlQM20    float64 `json:"BL_Q_M20,string" validate:"required"`
	BlQM21    float64 `json:"BL_Q_M21,string" validate:"required"`
	BlQM22    float64 `json:"BL_Q_M22,string" validate:"required"`
	BlQM23    float64 `json:"BL_Q_M23,string" validate:"required"`
	BlQM24    float64 `json:"BL_Q_M24,string" validate:"required"`
	NewBLM1   float64 `json:"New_BL_M1,string" validate:"required"`
	NewBLM2   float64 `json:"New_BL_M2,string" validate:"required"`
	NewBLM3   float64 `json:"New_BL_M3,string" validate:"required"`
	NewBLM4   float64 `json:"New_BL_M4,string" validate:"required"`
	NewBLM5   float64 `json:"New_BL_M5,string" validate:"required"`
	NewBLM6   float64 `json:"New_BL_M6,string" validate:"required"`
	NewBLM7   float64 `json:"New_BL_M7,string" validate:"required"`
	NewBLM8   float64 `json:"New_BL_M8,string" validate:"required"`
	NewBLM9   float64 `json:"New_BL_M9,string" validate:"required"`
	NewBLM10  float64 `json:"New_BL_M10,string" validate:"required"`
	NewBLM11  float64 `json:"New_BL_M11,string" validate:"required"`
	NewBLM12  float64 `json:"New_BL_M12,string" validate:"required"`
	NewBLM13  float64 `json:"New_BL_M13,string" validate:"required"`
	NewBLM14  float64 `json:"New_BL_M14,string" validate:"required"`
	NewBLM15  float64 `json:"New_BL_M15,string" validate:"required"`
	NewBLM16  float64 `json:"New_BL_M16,string" validate:"required"`
	NewBLM17  float64 `json:"New_BL_M17,string" validate:"required"`
	NewBLM18  float64 `json:"New_BL_M18,string" validate:"required"`
	NewBLM19  float64 `json:"New_BL_M19,string" validate:"required"`
	NewBLM20  float64 `json:"New_BL_M20,string" validate:"required"`
	NewBLM21  float64 `json:"New_BL_M21,string" validate:"required"`
	NewBLM22  float64 `json:"New_BL_M22,string" validate:"required"`
	NewBLM23  float64 `json:"New_BL_M23,string" validate:"required"`
	NewBLM24  float64 `json:"New_BL_M24,string" validate:"required"`
}
