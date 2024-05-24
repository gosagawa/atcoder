package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	defer flush()
	/*
		mp := make(map[int]bool)
		mx := 10000
		ts := []int{6, 10, 15}
		for i := 2; i <= 2500; i++ {
			for _, v := range ts {
				if v*i <= mx {
					mp[v*i] = true
				}
			}
		}
		rs := make([]int, len(mp))
		i := 0
		for v := range mp {
			rs[i] = v
			i++
		}
		sorti(rs)
		rs = append([]int{6, 10, 15}, rs...)
		outis(rs)
		return
	*/

	as := []int{6, 10, 15, 12, 18, 20, 24, 30, 36, 40, 42, 45, 48, 50, 54, 60, 66, 70, 72, 75, 78, 80, 84, 90, 96, 100, 102, 105, 108, 110, 114, 120, 126, 130, 132, 135, 138, 140, 144, 150, 156, 160, 162, 165, 168, 170, 174, 180, 186, 190, 192, 195, 198, 200, 204, 210, 216, 220, 222, 225, 228, 230, 234, 240, 246, 250, 252, 255, 258, 260, 264, 270, 276, 280, 282, 285, 288, 290, 294, 300, 306, 310, 312, 315, 318, 320, 324, 330, 336, 340, 342, 345, 348, 350, 354, 360, 366, 370, 372, 375, 378, 380, 384, 390, 396, 400, 402, 405, 408, 410, 414, 420, 426, 430, 432, 435, 438, 440, 444, 450, 456, 460, 462, 465, 468, 470, 474, 480, 486, 490, 492, 495, 498, 500, 504, 510, 516, 520, 522, 525, 528, 530, 534, 540, 546, 550, 552, 555, 558, 560, 564, 570, 576, 580, 582, 585, 588, 590, 594, 600, 606, 610, 612, 615, 618, 620, 624, 630, 636, 640, 642, 645, 648, 650, 654, 660, 666, 670, 672, 675, 678, 680, 684, 690, 696, 700, 702, 705, 708, 710, 714, 720, 726, 730, 732, 735, 738, 740, 744, 750, 756, 760, 762, 765, 768, 770, 774, 780, 786, 790, 792, 795, 798, 800, 804, 810, 816, 820, 822, 825, 828, 830, 834, 840, 846, 850, 852, 855, 858, 860, 864, 870, 876, 880, 882, 885, 888, 890, 894, 900, 906, 910, 912, 915, 918, 920, 924, 930, 936, 940, 942, 945, 948, 950, 954, 960, 966, 970, 972, 975, 978, 980, 984, 990, 996, 1000, 1002, 1005, 1008, 1010, 1014, 1020, 1026, 1030, 1032, 1035, 1038, 1040, 1044, 1050, 1056, 1060, 1062, 1065, 1068, 1070, 1074, 1080, 1086, 1090, 1092, 1095, 1098, 1100, 1104, 1110, 1116, 1120, 1122, 1125, 1128, 1130, 1134, 1140, 1146, 1150, 1152, 1155, 1158, 1160, 1164, 1170, 1176, 1180, 1182, 1185, 1188, 1190, 1194, 1200, 1206, 1210, 1212, 1215, 1218, 1220, 1224, 1230, 1236, 1240, 1242, 1245, 1248, 1250, 1254, 1260, 1266, 1270, 1272, 1275, 1278, 1280, 1284, 1290, 1296, 1300, 1302, 1305, 1308, 1310, 1314, 1320, 1326, 1330, 1332, 1335, 1338, 1340, 1344, 1350, 1356, 1360, 1362, 1365, 1368, 1370, 1374, 1380, 1386, 1390, 1392, 1395, 1398, 1400, 1404, 1410, 1416, 1420, 1422, 1425, 1428, 1430, 1434, 1440, 1446, 1450, 1452, 1455, 1458, 1460, 1464, 1470, 1476, 1480, 1482, 1485, 1488, 1490, 1494, 1500, 1506, 1510, 1512, 1515, 1518, 1520, 1524, 1530, 1536, 1540, 1542, 1545, 1548, 1550, 1554, 1560, 1566, 1570, 1572, 1575, 1578, 1580, 1584, 1590, 1596, 1600, 1602, 1605, 1608, 1610, 1614, 1620, 1626, 1630, 1632, 1635, 1638, 1640, 1644, 1650, 1656, 1660, 1662, 1665, 1668, 1670, 1674, 1680, 1686, 1690, 1692, 1695, 1698, 1700, 1704, 1710, 1716, 1720, 1722, 1725, 1728, 1730, 1734, 1740, 1746, 1750, 1752, 1755, 1758, 1760, 1764, 1770, 1776, 1780, 1782, 1785, 1788, 1790, 1794, 1800, 1806, 1810, 1812, 1815, 1818, 1820, 1824, 1830, 1836, 1840, 1842, 1845, 1848, 1850, 1854, 1860, 1866, 1870, 1872, 1875, 1878, 1880, 1884, 1890, 1896, 1900, 1902, 1905, 1908, 1910, 1914, 1920, 1926, 1930, 1932, 1935, 1938, 1940, 1944, 1950, 1956, 1960, 1962, 1965, 1968, 1970, 1974, 1980, 1986, 1990, 1992, 1995, 1998, 2000, 2004, 2010, 2016, 2020, 2022, 2025, 2028, 2030, 2034, 2040, 2046, 2050, 2052, 2055, 2058, 2060, 2064, 2070, 2076, 2080, 2082, 2085, 2088, 2090, 2094, 2100, 2106, 2110, 2112, 2115, 2118, 2120, 2124, 2130, 2136, 2140, 2142, 2145, 2148, 2150, 2154, 2160, 2166, 2170, 2172, 2175, 2178, 2180, 2184, 2190, 2196, 2200, 2202, 2205, 2208, 2210, 2214, 2220, 2226, 2230, 2232, 2235, 2238, 2240, 2244, 2250, 2256, 2260, 2262, 2265, 2268, 2270, 2274, 2280, 2286, 2290, 2292, 2295, 2298, 2300, 2304, 2310, 2316, 2320, 2322, 2325, 2328, 2330, 2334, 2340, 2346, 2350, 2352, 2355, 2358, 2360, 2364, 2370, 2376, 2380, 2382, 2385, 2388, 2390, 2394, 2400, 2406, 2410, 2412, 2415, 2418, 2420, 2424, 2430, 2436, 2440, 2442, 2445, 2448, 2450, 2454, 2460, 2466, 2470, 2472, 2475, 2478, 2480, 2484, 2490, 2496, 2500, 2502, 2505, 2508, 2510, 2514, 2520, 2526, 2530, 2532, 2535, 2538, 2540, 2544, 2550, 2556, 2560, 2562, 2565, 2568, 2570, 2574, 2580, 2586, 2590, 2592, 2595, 2598, 2600, 2604, 2610, 2616, 2620, 2622, 2625, 2628, 2630, 2634, 2640, 2646, 2650, 2652, 2655, 2658, 2660, 2664, 2670, 2676, 2680, 2682, 2685, 2688, 2690, 2694, 2700, 2706, 2710, 2712, 2715, 2718, 2720, 2724, 2730, 2736, 2740, 2742, 2745, 2748, 2750, 2754, 2760, 2766, 2770, 2772, 2775, 2778, 2780, 2784, 2790, 2796, 2800, 2802, 2805, 2808, 2810, 2814, 2820, 2826, 2830, 2832, 2835, 2838, 2840, 2844, 2850, 2856, 2860, 2862, 2865, 2868, 2870, 2874, 2880, 2886, 2890, 2892, 2895, 2898, 2900, 2904, 2910, 2916, 2920, 2922, 2925, 2928, 2930, 2934, 2940, 2946, 2950, 2952, 2955, 2958, 2960, 2964, 2970, 2976, 2980, 2982, 2985, 2988, 2990, 2994, 3000, 3006, 3010, 3012, 3015, 3018, 3020, 3024, 3030, 3036, 3040, 3042, 3045, 3048, 3050, 3054, 3060, 3066, 3070, 3072, 3075, 3078, 3080, 3084, 3090, 3096, 3100, 3102, 3105, 3108, 3110, 3114, 3120, 3126, 3130, 3132, 3135, 3138, 3140, 3144, 3150, 3156, 3160, 3162, 3165, 3168, 3170, 3174, 3180, 3186, 3190, 3192, 3195, 3198, 3200, 3204, 3210, 3216, 3220, 3222, 3225, 3228, 3230, 3234, 3240, 3246, 3250, 3252, 3255, 3258, 3260, 3264, 3270, 3276, 3280, 3282, 3285, 3288, 3290, 3294, 3300, 3306, 3310, 3312, 3315, 3318, 3320, 3324, 3330, 3336, 3340, 3342, 3345, 3348, 3350, 3354, 3360, 3366, 3370, 3372, 3375, 3378, 3380, 3384, 3390, 3396, 3400, 3402, 3405, 3408, 3410, 3414, 3420, 3426, 3430, 3432, 3435, 3438, 3440, 3444, 3450, 3456, 3460, 3462, 3465, 3468, 3470, 3474, 3480, 3486, 3490, 3492, 3495, 3498, 3500, 3504, 3510, 3516, 3520, 3522, 3525, 3528, 3530, 3534, 3540, 3546, 3550, 3552, 3555, 3558, 3560, 3564, 3570, 3576, 3580, 3582, 3585, 3588, 3590, 3594, 3600, 3606, 3610, 3612, 3615, 3618, 3620, 3624, 3630, 3636, 3640, 3642, 3645, 3648, 3650, 3654, 3660, 3666, 3670, 3672, 3675, 3678, 3680, 3684, 3690, 3696, 3700, 3702, 3705, 3708, 3710, 3714, 3720, 3726, 3730, 3732, 3735, 3738, 3740, 3744, 3750, 3756, 3760, 3762, 3765, 3768, 3770, 3774, 3780, 3786, 3790, 3792, 3795, 3798, 3800, 3804, 3810, 3816, 3820, 3822, 3825, 3828, 3830, 3834, 3840, 3846, 3850, 3852, 3855, 3858, 3860, 3864, 3870, 3876, 3880, 3882, 3885, 3888, 3890, 3894, 3900, 3906, 3910, 3912, 3915, 3918, 3920, 3924, 3930, 3936, 3940, 3942, 3945, 3948, 3950, 3954, 3960, 3966, 3970, 3972, 3975, 3978, 3980, 3984, 3990, 3996, 4000, 4002, 4005, 4008, 4010, 4014, 4020, 4026, 4030, 4032, 4035, 4038, 4040, 4044, 4050, 4056, 4060, 4062, 4065, 4068, 4070, 4074, 4080, 4086, 4090, 4092, 4095, 4098, 4100, 4104, 4110, 4116, 4120, 4122, 4125, 4128, 4130, 4134, 4140, 4146, 4150, 4152, 4155, 4158, 4160, 4164, 4170, 4176, 4180, 4182, 4185, 4188, 4190, 4194, 4200, 4206, 4210, 4212, 4215, 4218, 4220, 4224, 4230, 4236, 4240, 4242, 4245, 4248, 4250, 4254, 4260, 4266, 4270, 4272, 4275, 4278, 4280, 4284, 4290, 4296, 4300, 4302, 4305, 4308, 4310, 4314, 4320, 4326, 4330, 4332, 4335, 4338, 4340, 4344, 4350, 4356, 4360, 4362, 4365, 4368, 4370, 4374, 4380, 4386, 4390, 4392, 4395, 4398, 4400, 4404, 4410, 4416, 4420, 4422, 4425, 4428, 4430, 4434, 4440, 4446, 4450, 4452, 4455, 4458, 4460, 4464, 4470, 4476, 4480, 4482, 4485, 4488, 4490, 4494, 4500, 4506, 4510, 4512, 4515, 4518, 4520, 4524, 4530, 4536, 4540, 4542, 4545, 4548, 4550, 4554, 4560, 4566, 4570, 4572, 4575, 4578, 4580, 4584, 4590, 4596, 4600, 4602, 4605, 4608, 4610, 4614, 4620, 4626, 4630, 4632, 4635, 4638, 4640, 4644, 4650, 4656, 4660, 4662, 4665, 4668, 4670, 4674, 4680, 4686, 4690, 4692, 4695, 4698, 4700, 4704, 4710, 4716, 4720, 4722, 4725, 4728, 4730, 4734, 4740, 4746, 4750, 4752, 4755, 4758, 4760, 4764, 4770, 4776, 4780, 4782, 4785, 4788, 4790, 4794, 4800, 4806, 4810, 4812, 4815, 4818, 4820, 4824, 4830, 4836, 4840, 4842, 4845, 4848, 4850, 4854, 4860, 4866, 4870, 4872, 4875, 4878, 4880, 4884, 4890, 4896, 4900, 4902, 4905, 4908, 4910, 4914, 4920, 4926, 4930, 4932, 4935, 4938, 4940, 4944, 4950, 4956, 4960, 4962, 4965, 4968, 4970, 4974, 4980, 4986, 4990, 4992, 4995, 4998, 5000, 5004, 5010, 5016, 5020, 5022, 5025, 5028, 5030, 5034, 5040, 5046, 5050, 5052, 5055, 5058, 5060, 5064, 5070, 5076, 5080, 5082, 5085, 5088, 5090, 5094, 5100, 5106, 5110, 5112, 5115, 5118, 5120, 5124, 5130, 5136, 5140, 5142, 5145, 5148, 5150, 5154, 5160, 5166, 5170, 5172, 5175, 5178, 5180, 5184, 5190, 5196, 5200, 5202, 5205, 5208, 5210, 5214, 5220, 5226, 5230, 5232, 5235, 5238, 5240, 5244, 5250, 5256, 5260, 5262, 5265, 5268, 5270, 5274, 5280, 5286, 5290, 5292, 5295, 5298, 5300, 5304, 5310, 5316, 5320, 5322, 5325, 5328, 5330, 5334, 5340, 5346, 5350, 5352, 5355, 5358, 5360, 5364, 5370, 5376, 5380, 5382, 5385, 5388, 5390, 5394, 5400, 5406, 5410, 5412, 5415, 5418, 5420, 5424, 5430, 5436, 5440, 5442, 5445, 5448, 5450, 5454, 5460, 5466, 5470, 5472, 5475, 5478, 5480, 5484, 5490, 5496, 5500, 5502, 5505, 5508, 5510, 5514, 5520, 5526, 5530, 5532, 5535, 5538, 5540, 5544, 5550, 5556, 5560, 5562, 5565, 5568, 5570, 5574, 5580, 5586, 5590, 5592, 5595, 5598, 5600, 5604, 5610, 5616, 5620, 5622, 5625, 5628, 5630, 5634, 5640, 5646, 5650, 5652, 5655, 5658, 5660, 5664, 5670, 5676, 5680, 5682, 5685, 5688, 5690, 5694, 5700, 5706, 5710, 5712, 5715, 5718, 5720, 5724, 5730, 5736, 5740, 5742, 5745, 5748, 5750, 5754, 5760, 5766, 5770, 5772, 5775, 5778, 5780, 5784, 5790, 5796, 5800, 5802, 5805, 5808, 5810, 5814, 5820, 5826, 5830, 5832, 5835, 5838, 5840, 5844, 5850, 5856, 5860, 5862, 5865, 5868, 5870, 5874, 5880, 5886, 5890, 5892, 5895, 5898, 5900, 5904, 5910, 5916, 5920, 5922, 5925, 5928, 5930, 5934, 5940, 5946, 5950, 5952, 5955, 5958, 5960, 5964, 5970, 5976, 5980, 5982, 5985, 5988, 5990, 5994, 6000, 6006, 6010, 6012, 6015, 6018, 6020, 6024, 6030, 6036, 6040, 6042, 6045, 6048, 6050, 6054, 6060, 6066, 6070, 6072, 6075, 6078, 6080, 6084, 6090, 6096, 6100, 6102, 6105, 6108, 6110, 6114, 6120, 6126, 6130, 6132, 6135, 6138, 6140, 6144, 6150, 6156, 6160, 6162, 6165, 6168, 6170, 6174, 6180, 6186, 6190, 6192, 6195, 6198, 6200, 6204, 6210, 6216, 6220, 6222, 6225, 6228, 6230, 6234, 6240, 6246, 6250, 6252, 6255, 6258, 6260, 6264, 6270, 6276, 6280, 6282, 6285, 6288, 6290, 6294, 6300, 6306, 6310, 6312, 6315, 6318, 6320, 6324, 6330, 6336, 6340, 6342, 6345, 6348, 6350, 6354, 6360, 6366, 6370, 6372, 6375, 6378, 6380, 6384, 6390, 6396, 6400, 6402, 6405, 6408, 6410, 6414, 6420, 6426, 6430, 6432, 6435, 6438, 6440, 6444, 6450, 6456, 6460, 6462, 6465, 6468, 6470, 6474, 6480, 6486, 6490, 6492, 6495, 6498, 6500, 6504, 6510, 6516, 6520, 6522, 6525, 6528, 6530, 6534, 6540, 6546, 6550, 6552, 6555, 6558, 6560, 6564, 6570, 6576, 6580, 6582, 6585, 6588, 6590, 6594, 6600, 6606, 6610, 6612, 6615, 6618, 6620, 6624, 6630, 6636, 6640, 6642, 6645, 6648, 6650, 6654, 6660, 6666, 6670, 6672, 6675, 6678, 6680, 6684, 6690, 6696, 6700, 6702, 6705, 6708, 6710, 6714, 6720, 6726, 6730, 6732, 6735, 6738, 6740, 6744, 6750, 6756, 6760, 6762, 6765, 6768, 6770, 6774, 6780, 6786, 6790, 6792, 6795, 6798, 6800, 6804, 6810, 6816, 6820, 6822, 6825, 6828, 6830, 6834, 6840, 6846, 6850, 6852, 6855, 6858, 6860, 6864, 6870, 6876, 6880, 6882, 6885, 6888, 6890, 6894, 6900, 6906, 6910, 6912, 6915, 6918, 6920, 6924, 6930, 6936, 6940, 6942, 6945, 6948, 6950, 6954, 6960, 6966, 6970, 6972, 6975, 6978, 6980, 6984, 6990, 6996, 7000, 7002, 7005, 7008, 7010, 7014, 7020, 7026, 7030, 7032, 7035, 7038, 7040, 7044, 7050, 7056, 7060, 7062, 7065, 7068, 7070, 7074, 7080, 7086, 7090, 7092, 7095, 7098, 7100, 7104, 7110, 7116, 7120, 7122, 7125, 7128, 7130, 7134, 7140, 7146, 7150, 7152, 7155, 7158, 7160, 7164, 7170, 7176, 7180, 7182, 7185, 7188, 7190, 7194, 7200, 7206, 7210, 7212, 7215, 7218, 7220, 7224, 7230, 7236, 7240, 7242, 7245, 7248, 7250, 7254, 7260, 7266, 7270, 7272, 7275, 7278, 7280, 7284, 7290, 7296, 7300, 7302, 7305, 7308, 7310, 7314, 7320, 7326, 7330, 7332, 7335, 7338, 7340, 7344, 7350, 7356, 7360, 7362, 7365, 7368, 7370, 7374, 7380, 7386, 7390, 7392, 7395, 7398, 7400, 7404, 7410, 7416, 7420, 7422, 7425, 7428, 7430, 7434, 7440, 7446, 7450, 7452, 7455, 7458, 7460, 7464, 7470, 7476, 7480, 7482, 7485, 7488, 7490, 7494, 7500, 7506, 7510, 7512, 7515, 7518, 7520, 7524, 7530, 7536, 7540, 7542, 7545, 7548, 7550, 7554, 7560, 7566, 7570, 7572, 7575, 7578, 7580, 7584, 7590, 7596, 7600, 7602, 7605, 7608, 7610, 7614, 7620, 7626, 7630, 7632, 7635, 7638, 7640, 7644, 7650, 7656, 7660, 7662, 7665, 7668, 7670, 7674, 7680, 7686, 7690, 7692, 7695, 7698, 7700, 7704, 7710, 7716, 7720, 7722, 7725, 7728, 7730, 7734, 7740, 7746, 7750, 7752, 7755, 7758, 7760, 7764, 7770, 7776, 7780, 7782, 7785, 7788, 7790, 7794, 7800, 7806, 7810, 7812, 7815, 7818, 7820, 7824, 7830, 7836, 7840, 7842, 7845, 7848, 7850, 7854, 7860, 7866, 7870, 7872, 7875, 7878, 7880, 7884, 7890, 7896, 7900, 7902, 7905, 7908, 7910, 7914, 7920, 7926, 7930, 7932, 7935, 7938, 7940, 7944, 7950, 7956, 7960, 7962, 7965, 7968, 7970, 7974, 7980, 7986, 7990, 7992, 7995, 7998, 8000, 8004, 8010, 8016, 8020, 8022, 8025, 8028, 8030, 8034, 8040, 8046, 8050, 8052, 8055, 8058, 8060, 8064, 8070, 8076, 8080, 8082, 8085, 8088, 8090, 8094, 8100, 8106, 8110, 8112, 8115, 8118, 8120, 8124, 8130, 8136, 8140, 8142, 8145, 8148, 8150, 8154, 8160, 8166, 8170, 8172, 8175, 8178, 8180, 8184, 8190, 8196, 8200, 8202, 8205, 8208, 8210, 8214, 8220, 8226, 8230, 8232, 8235, 8238, 8240, 8244, 8250, 8256, 8260, 8262, 8265, 8268, 8270, 8274, 8280, 8286, 8290, 8292, 8295, 8298, 8300, 8304, 8310, 8316, 8320, 8322, 8325, 8328, 8330, 8334, 8340, 8346, 8350, 8352, 8355, 8358, 8360, 8364, 8370, 8376, 8380, 8382, 8385, 8388, 8390, 8394, 8400, 8406, 8410, 8412, 8415, 8418, 8420, 8424, 8430, 8436, 8440, 8442, 8445, 8448, 8450, 8454, 8460, 8466, 8470, 8472, 8475, 8478, 8480, 8484, 8490, 8496, 8500, 8502, 8505, 8508, 8510, 8514, 8520, 8526, 8530, 8532, 8535, 8538, 8540, 8544, 8550, 8556, 8560, 8562, 8565, 8568, 8570, 8574, 8580, 8586, 8590, 8592, 8595, 8598, 8600, 8604, 8610, 8616, 8620, 8622, 8625, 8628, 8630, 8634, 8640, 8646, 8650, 8652, 8655, 8658, 8660, 8664, 8670, 8676, 8680, 8682, 8685, 8688, 8690, 8694, 8700, 8706, 8710, 8712, 8715, 8718, 8720, 8724, 8730, 8736, 8740, 8742, 8745, 8748, 8750, 8754, 8760, 8766, 8770, 8772, 8775, 8778, 8780, 8784, 8790, 8796, 8800, 8802, 8805, 8808, 8810, 8814, 8820, 8826, 8830, 8832, 8835, 8838, 8840, 8844, 8850, 8856, 8860, 8862, 8865, 8868, 8870, 8874, 8880, 8886, 8890, 8892, 8895, 8898, 8900, 8904, 8910, 8916, 8920, 8922, 8925, 8928, 8930, 8934, 8940, 8946, 8950, 8952, 8955, 8958, 8960, 8964, 8970, 8976, 8980, 8982, 8985, 8988, 8990, 8994, 9000, 9006, 9010, 9012, 9015, 9018, 9020, 9024, 9030, 9036, 9040, 9042, 9045, 9048, 9050, 9054, 9060, 9066, 9070, 9072, 9075, 9078, 9080, 9084, 9090, 9096, 9100, 9102, 9105, 9108, 9110, 9114, 9120, 9126, 9130, 9132, 9135, 9138, 9140, 9144, 9150, 9156, 9160, 9162, 9165, 9168, 9170, 9174, 9180, 9186, 9190, 9192, 9195, 9198, 9200, 9204, 9210, 9216, 9220, 9222, 9225, 9228, 9230, 9234, 9240, 9246, 9250, 9252, 9255, 9258, 9260, 9264, 9270, 9276, 9280, 9282, 9285, 9288, 9290, 9294, 9300, 9306, 9310, 9312, 9315, 9318, 9320, 9324, 9330, 9336, 9340, 9342, 9345, 9348, 9350, 9354, 9360, 9366, 9370, 9372, 9375, 9378, 9380, 9384, 9390, 9396, 9400, 9402, 9405, 9408, 9410, 9414, 9420, 9426, 9430, 9432, 9435, 9438, 9440, 9444, 9450, 9456, 9460, 9462, 9465, 9468, 9470, 9474, 9480, 9486, 9490, 9492, 9495, 9498, 9500, 9504, 9510, 9516, 9520, 9522, 9525, 9528, 9530, 9534, 9540, 9546, 9550, 9552, 9555, 9558, 9560, 9564, 9570, 9576, 9580, 9582, 9585, 9588, 9590, 9594, 9600, 9606, 9610, 9612, 9615, 9618, 9620, 9624, 9630, 9636, 9640, 9642, 9645, 9648, 9650, 9654, 9660, 9666, 9670, 9672, 9675, 9678, 9680, 9684, 9690, 9696, 9700, 9702, 9705, 9708, 9710, 9714, 9720, 9726, 9730, 9732, 9735, 9738, 9740, 9744, 9750, 9756, 9760, 9762, 9765, 9768, 9770, 9774, 9780, 9786, 9790, 9792, 9795, 9798, 9800, 9804, 9810, 9816, 9820, 9822, 9825, 9828, 9830, 9834, 9840, 9846, 9850, 9852, 9855, 9858, 9860, 9864, 9870, 9876, 9880, 9882, 9885, 9888, 9890, 9894, 9900, 9906, 9910, 9912, 9915, 9918, 9920, 9924, 9930, 9936, 9940, 9942, 9945, 9948, 9950, 9954, 9960, 9966, 9970, 9972, 9975, 9978, 9980, 9984, 9990, 9996, 10000}

	n := ni()
	outis(as[:n])
}

// ==================================================
// init
// ==================================================

const inf = math.MaxInt64
const mod1000000007 = 1000000007
const mod998244353 = 998244353
const mod = mod1000000007
const baseRune = 'a'
const maxlogn = 62

var mpowcache map[[3]int]int
var debugFlg bool

func init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
		debugFlg = true
	}
	mpowcache = make(map[[3]int]int)
}

// ==================================================
// io
// ==================================================

func ni() int {
	sc.Scan()
	return atoi(sc.Text())
}

func ni2() (int, int) {
	return ni(), ni()
}

func ni3() (int, int, int) {
	return ni(), ni(), ni()
}

func ni4() (int, int, int, int) {
	return ni(), ni(), ni(), ni()
}

func nis(arg ...int) []int {
	n := arg[0]
	t := 0
	if len(arg) == 2 {
		t = arg[1]
	}

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = ni() - t
	}
	return a
}

func ni2s(n int) ([]int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = ni2()
	}
	return a, b
}

func ni3s(n int) ([]int, []int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i] = ni3()
	}
	return a, b, c
}

func ni4s(n int) ([]int, []int, []int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i], c[i], d[i] = ni4()
	}
	return a, b, c, d
}

func ni2a(n int) [][2]int {
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		a[i][0], a[i][1] = ni2()
	}
	return a
}

func ni3a(n int) [][3]int {
	a := make([][3]int, n)
	for i := 0; i < n; i++ {
		a[i][0], a[i][1], a[i][2] = ni3()
	}
	return a
}

func ni4a(n int) [][4]int {
	a := make([][4]int, n)
	for i := 0; i < n; i++ {
		a[i][0], a[i][1], a[i][2], a[i][3] = ni4()
	}
	return a
}

func ni2d(n, m int) [][]int {
	a := i2s(n, m, 0)
	for i := 0; i < n; i++ {
		a[i] = nis(m)
	}
	return a
}

func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func nsis() []int {
	sc.Scan()
	s := sc.Text()
	return stois(s, baseRune)
}

func nsi2s(n int) [][]int {
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = nsis()
	}
	return mp
}

// mp := convidxi2s(nsi2s(n), map[string]int{".": 0, "#": 1})
func convidxi2s(sl [][]int, conv map[string]int) [][]int {
	imap := make(map[int]int)
	for s, v := range conv {
		imap[ctoi(s)] = v
	}
	for i, sl2 := range sl {
		for j, v := range sl2 {
			sl[i][j] = imap[v]
		}
	}
	return sl
}

// mp := convidxis(nsis(), map[string]int{".": 0, "#": 1})
func convidxis(sl []int, conv map[string]int) []int {
	imap := make(map[int]int)
	for s, v := range conv {
		imap[ctoi(s)] = v
	}
	for i, v := range sl {
		sl[i] = imap[v]
	}
	return sl
}

func ctoi(c string) int {
	return int(rune(c[0]) - baseRune)
}

func nsiis() []int {
	sc.Scan()
	s := sc.Text()
	return stois(s, '0')
}

func scani() int {
	var i int
	fmt.Scanf("%i", &i)
	return i
}

func scans() string {
	var s string
	fmt.Scanf("%s", &s)
	return s
}

func out(v ...interface{}) {
	_, e := fmt.Fprintln(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func debug(v ...interface{}) {
	if !debugFlg {
		return
	}
	out(v...)
}

func debugi2s(sl [][]int) {
	if !debugFlg {
		return
	}
	for _, v := range sl {
		outis(v)
	}
	out("")
}

func outf(f string, v ...interface{}) {
	out(fmt.Sprintf(f, v...))
}

func outwoln(v ...interface{}) {
	_, e := fmt.Fprint(wtr, v...)
	if e != nil {
		panic(e)
	}
}

func outis(sl []int) {
	r := make([]string, len(sl))
	for i, v := range sl {
		r[i] = itoa(v)
	}
	out(strings.Join(r, " "))
}

func outisnr(sl []int) {
	for _, v := range sl {
		out(v)
	}
}

func out2d(i, j int) {
	outf("%v %v", i, j)
}

func outsj(sl []string) {
	out(sj(sl))
}

func outsjsp(sl []string) {
	out(sjsp(sl))
}

func outfl(v float64) {
	outf("%.15f", v)
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}

func nftoi(decimalLen int) int {
	sc.Scan()
	s := sc.Text()

	r := 0
	minus := strings.Split(s, "-")
	isMinus := false
	if len(minus) == 2 {
		s = minus[1]
		isMinus = true
	}

	t := strings.Split(s, ".")
	i := atoi(t[0])
	r += i * pow(10, decimalLen)
	if len(t) > 1 {
		i = atoi(t[1])
		i *= pow(10, decimalLen-len(t[1]))
		r += i
	}
	if isMinus {
		return -r
	}
	return r
}

func atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func bytoi(b byte) int {
	return atoi(string(b))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// ==================================================
// num
// ==================================================

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxs(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mins(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func maxf(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func maxsf(a *float64, b float64) {
	if *a < b {
		*a = b
	}
}

func minf(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func minsf(a *float64, b float64) {
	if *a > b {
		*a = b
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

var pow2cache [64]int

func pow2(i int) int {
	if pow2cache[i] == 0 {
		pow2cache[i] = int(math.Pow(2, float64(i)))
	}
	return pow2cache[i]
}

var pow10cache [20]int

func pow10(i int) int {
	if pow10cache[i] == 0 {
		pow10cache[i] = int(math.Pow(10, float64(i)))
	}
	return pow10cache[i]
}

func sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func sqrtf(i int) float64 {
	return math.Sqrt(float64(i))
}

func ch(cond bool, ok, ng int) int {
	if cond {
		return ok
	}
	return ng
}

func mul(a, b int) (int, int) {
	if a < 0 {
		a, b = -a, -b
	}
	if a == 0 || b == 0 {
		return 0, 0
	} else if a > 0 && b > 0 && a > math.MaxInt64/b {
		return 0, +1
	} else if a < math.MinInt64/b {
		return 0, -1
	}
	return a * b, 0
}

func getAngle(x, y float64) float64 {
	return math.Atan2(y, x) * 180 / math.Pi
}

func permutation(n int, k int) int {
	if k > n || k <= 0 {
		panic(fmt.Sprintf("invalid param n:%v k:%v", n, k))
	}
	v := 1
	for i := 0; i < k; i++ {
		v *= (n - i)
	}
	return v
}

/*
	for {

		// Do something

		if !nextPermutation(sort.IntSlice(x)) {
			break
		}
	}
*/
func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

type combFactorial struct {
	fac    []int
	facinv []int
}

func newcombFactorial(n int) *combFactorial {

	fac := make([]int, n)
	facinv := make([]int, n)
	fac[0] = 1
	facinv[0] = minvfermat(1, mod)

	for i := 1; i < n; i++ {
		fac[i] = mmul(i, fac[i-1])
		facinv[i] = minvfermat(fac[i], mod)
	}

	return &combFactorial{
		fac:    fac,
		facinv: facinv,
	}
}

func (c *combFactorial) factorial(n int) int {
	return c.fac[n]
}

func (c *combFactorial) combination(n, r int) int {
	if r > n {
		return 0
	}
	return mmul(mmul(c.fac[n], c.facinv[r]), c.facinv[n-r])
}

func (c *combFactorial) permutation(n, r int) int {
	if r > n {
		return 0
	}
	return mmul(c.fac[n], c.facinv[n-r])
}

func (c *combFactorial) homogeousProduct(n, r int) int {
	return c.combination(n-1+r, r)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func gcm(a, b int) int {
	t := gcd(a, b)
	return a * b / t
}

func divisor(n int) ([]int, map[int]int) {
	sqrtn := int(math.Sqrt(float64(n)))
	c := 2
	divisor := []int{}
	divisorm := make(map[int]int)
	for {
		if n%2 != 0 {
			break
		}
		divisor = append(divisor, 2)
		divisorm[2]++
		n /= 2
	}
	c = 3
	for {
		if n%c == 0 {
			divisor = append(divisor, c)
			divisorm[c]++
			n /= c
		} else {
			c += 2
			if c > sqrtn {
				break
			}
		}
	}
	if n != 1 {
		divisor = append(divisor, n)
		divisorm[n]++
	}
	return divisor, divisorm
}

func alldivisor(n int) []int {
	sqrtn := int(math.Sqrt(float64(n)))
	divisor := []int{}
	for i := 1; i <= sqrtn; i++ {
		if n%i != 0 {
			continue
		}
		divisor = append(divisor, i)
		if n/i != i {
			divisor = append(divisor, n/i)
		}
	}
	return divisor
}

func mmod(a, m int) int {
	return (a%m + m) % m
}

func extGcd(a, b int) (int, int, int) {
	return extGcdSub(b, a%b, 0, 0)
}

func extGcdSub(a, b, p, q int) (int, int, int) {
	if b == 0 {
		return 1, 0, a
	}
	q, p, d := extGcdSub(b, a%b, q, p)
	q -= a / b * p
	return p, q, d
}

func chineseRem(b1, m1, b2, m2 int) (bool, int, int) {
	p, _, d := extGcd(m1, m2)
	if (b2-b1)%d != 0 {
		return false, 0, 0
	}
	m := m1 * (m2 / d)
	tmp := (b2 - b1) / d * p % (m2 / d)
	r := mmod(b1+m1*tmp, m)
	return true, r, m
}

type binom struct {
	fac  []int
	finv []int
	inv  []int
}

func newbinom(n int) *binom {
	b := &binom{
		fac:  make([]int, n),
		finv: make([]int, n),
		inv:  make([]int, n),
	}
	b.fac[0] = 1
	b.fac[1] = 1
	b.inv[1] = 1
	b.finv[0] = 1
	b.finv[1] = 1
	for i := 2; i < n; i++ {
		b.fac[i] = b.fac[i-1] * i % mod
		b.inv[i] = mod - mod/i*b.inv[mod%i]%mod
		b.finv[i] = b.finv[i-1] * b.inv[i] % mod
	}
	return b
}

func (b *binom) get(n, r int) int {
	if n < r || n < 0 || r < 0 {
		return 0
	}
	return b.fac[n] * b.finv[r] % mod * b.finv[n-r] % mod
}

func matPow(a [][]int, n int) [][]int {
	r := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = is(len(a), 0)
		r[i][i] = 1
	}
	for n > 0 {
		if n&1 != 0 {
			r = matMul(a, r)
		}
		a = matMul(a, a)
		n = n >> 1
	}
	return r
}

func matMul(a, b [][]int) [][]int {
	r := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		r[i] = is(len(b[0]), 0)
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				r[i][j] = madd(r[i][j], mmul(a[i][k], b[k][j]))
			}
		}
	}
	return r
}

// ==================================================
// mod
// ==================================================

func madd(a, b int) int {
	a %= mod
	b %= mod
	a += b
	if a >= mod || a <= -mod {
		a %= mod
	}
	if a < 0 {
		a += mod
	}
	return a
}

func mmul(a, b int) int {
	a %= mod
	b %= mod
	return a * b % mod
}

func mdiv(a, b int) int {
	a %= mod
	if b <= 0 || b >= mod {
		panic("invalid division")
	}
	return a * minvfermat(b, mod) % mod
}

func mpow(a, n, m int) int {
	if v, ok := mpowcache[[3]int{a, n, m}]; ok {
		return v
	}
	fa := a
	fn := n
	if m == 1 {
		return 0
	}
	r := 1
	for n > 0 {
		if n&1 == 1 {
			r = r * a % m
		}
		a, n = a*a%m, n>>1
	}
	mpowcache[[3]int{fa, fn, m}] = r
	return r
}

func minv(a, m int) int {
	p, x, u := m, 1, 0
	for p != 0 {
		t := a / p
		a, p = p, a-t*p
		x, u = u, x-t*u
	}
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

// m only allow prime number
func minvfermat(a, m int) int {
	return mpow(a, m-2, mod)
}

// ==================================================
// binarysearch
// ==================================================

/*
	o = bs(0, len(sl)-1, func(c int) bool {
		return true
	})
*/
func bs(ok, ng int, f func(int) bool) int {
	if !f(ok) {
		return -1
	}
	if f(ng) {
		return ng
	}
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

/*
	o = bsfl(0.0, 100.0, 100, func(c float64) bool {
		return true
	})
*/
func bsfl(ok, ng float64, c int, f func(float64) bool) float64 {
	for i := 0; i < c; i++ {

		mid := (ok + ng) / 2

		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func bs3fl(low, high float64, c int, f func(float64) float64) float64 {

	for i := 0; i < c; i++ {
		c1 := (low*2 + high) / 3
		c2 := (low + high*2) / 3

		if f(c1) > f(c2) {
			low = c1
		} else {
			high = c2
		}
	}
	return low
}

func bs3i(low, high int, f func(int) int) (int, int) {

	for high-low > 2 {
		c1 := (low*2 + high) / 3
		c2 := (low + high*2) / 3
		fc1 := f(c1)
		fc2 := f(c2)

		if fc1 > fc2 {
			low = c1
		} else {
			high = c2
		}
	}
	if high-low == 2 {
		fc1 := f(low)
		fc2 := f(high)
		//	out(high, low, c1, c2, fc1, fc2)

		if fc1 > fc2 {
			low++
		} else {
			high--
		}
	}
	ri := 0
	rv := 0
	if high-low == 1 {
		fc1 := f(low)
		fc2 := f(high)
		//	out(high, low, c1, c2, fc1, fc2)

		if fc1 > fc2 {
			ri = high
			rv = fc2
		} else {
			ri = low
			rv = fc1
		}
	}
	return ri, rv
}

// ==================================================
// bit
// ==================================================

func hasbit(a int, n int) bool {
	return (a>>uint(n))&1 == 1
}

func nthbit(a int, n int) int {
	return int((a >> uint(n)) & 1)
}

func popcount(a int) int {
	return bits.OnesCount(uint(a))
}

func bitlen(a int) int {
	return bits.Len(uint(a))
}

func xor(a, b bool) bool { return a != b }

func debugbit(n int) string {
	r := ""
	for i := bitlen(n) - 1; i >= 0; i-- {
		if n&(1<<i) != 0 {
			r += "1"
		} else {
			r += "0"
		}
	}
	return r
}

// ==================================================
// string
// ==================================================

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func isLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func isUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

// ==================================================
// sort
// ==================================================

type sortOrder int

const (
	asc sortOrder = iota
	desc
)

func sorti(sl []int) {
	sort.Sort(sort.IntSlice(sl))
}

func sortir(sl []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
}

func sorts(sl []string) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
}

type Sort2ArOptions struct {
	keys   []int
	orders []sortOrder
}

type Sort2ArOption func(*Sort2ArOptions)

func opt2ar(key int, order sortOrder) Sort2ArOption {
	return func(args *Sort2ArOptions) {
		args.keys = append(args.keys, key)
		args.orders = append(args.orders, order)
	}
}

// sort2ar(sl,opt2ar(1,asc))
// sort2ar(sl,opt2ar(0,asc),opt2ar(1,asc))
func sort2ar(sl [][2]int, setters ...Sort2ArOption) {
	args := &Sort2ArOptions{}

	for _, setter := range setters {
		setter(args)
	}

	sort.Slice(sl, func(i, j int) bool {
		for idx, key := range args.keys {
			if sl[i][key] == sl[j][key] {
				continue
			}
			switch args.orders[idx] {
			case asc:
				return sl[i][key] < sl[j][key]
			case desc:
				return sl[i][key] > sl[j][key]
			}
		}
		return true
	})
}

// ==================================================
// slice
// ==================================================

func is(l int, def int) []int {
	sl := make([]int, l)
	for i := 0; i < l; i++ {
		sl[i] = def
	}
	return sl
}

func i2s(l, m int, def int) [][]int {
	sl := make([][]int, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}

func ss(l int) []string {
	return make([]string, l)
}

func sj(sl []string) string {
	return strings.Join(sl, "")
}

func sjsp(sl []string) string {
	return strings.Join(sl, " ")
}

//	out(stois("abcde", 'a'))
//	out(stois("abcde", 'a'-1))
//	out(stois("12345", '0'))
func stois(s string, baseRune rune) []int {
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(v - baseRune)
	}
	return r
}

func istos(s []int, baseRune rune) string {
	r := make([]byte, len(s))
	for i, v := range s {
		r[i] = byte(v) + byte(baseRune)
	}
	return string(r)
}

func issum(sl []int) int {
	r := 0
	for _, v := range sl {
		r += v
	}
	return r
}

func issummod(sl []int) int {
	r := 0
	for _, v := range sl {
		r += v
		r %= mod
	}
	return r
}

func reverse(sl []interface{}) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func reversei(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func uniquei(sl []int) []int {
	hist := make(map[int]struct{})
	j := 0
	rsl := make([]int, len(sl))
	for i := 0; i < len(sl); i++ {
		if _, ok := hist[sl[i]]; ok {
			continue
		}
		rsl[j] = sl[i]
		hist[sl[i]] = struct{}{}
		j++
	}
	return rsl[:j]
}

// coordinate compression
func cocom(sl []int) ([]int, map[int]int) {
	rsl := uniquei(sl)
	sorti(rsl)
	rm := make(map[int]int)
	for i := 0; i < len(rsl); i++ {
		rm[rsl[i]] = i
	}
	return rsl, rm
}

func popBack(sl []int) (int, []int) {
	return sl[len(sl)-1], sl[:len(sl)-1]
}

func addIdx(pos, v int, sl []int) []int {
	if len(sl) == pos {
		sl = append(sl, v)
		return sl
	}
	sl = append(sl[:pos+1], sl[pos:]...)
	sl[pos] = v
	return sl
}

func delIdx(pos int, sl []int) []int {
	return append(sl[:pos], sl[pos+1:]...)
}

// find x of sl[x] < v. return -1 if no lowerbound found
func lowerBound(v int, sl []int) int {
	if len(sl) == 0 {
		return -1
	}
	idx := bs(0, len(sl)-1, func(c int) bool {
		return sl[c] < v
	})
	return idx
}

// find x of v < sl[x]. return len(sl) if no upperbound found
func upperBound(v int, sl []int) int {
	if len(sl) == 0 {
		return 0
	}
	idx := bs(0, len(sl)-1, func(c int) bool {
		return sl[c] <= v
	})
	return idx + 1
}

func rotate(sl [][]int) [][]int {
	n := len(sl)
	r := i2s(n, n, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r[i][j] = sl[n-1-j][i]
		}
	}
	return r
}

// ==================================================
// matrix
// ==================================================

func matrixmul(a, b [][]int) [][]int {
	ac := len(a)
	ar := len(a[0])
	bc := len(b)
	br := len(b[0])
	if ar != bc {
		panic(fmt.Sprintf("invalid matrix mul ar:%v bc:%v", ar, bc))
	}

	r := i2s(ac, br, 0)
	for i := 0; i < ac; i++ {
		for j := 0; j < br; j++ {
			for k := 0; k < ar; k++ {
				r[i][j] += mmul(a[i][k], b[k][j])
				r[i][j] %= mod
			}
		}
	}
	return r
}

func slmatrixmul(a []int, b [][]int) []int {
	ar := len(a)
	bc := len(b)
	br := len(b[0])
	if ar != bc {
		panic(fmt.Sprintf("invalid matrix mul ar:%v bc:%v", ar, bc))
	}
	r := is(br, 0)
	for i := 0; i < br; i++ {
		for j := 0; j < ar; j++ {
			r[i] += mmul(a[j], b[j][i])
			r[i] %= mod
		}
	}
	return r
}

func matrixpow(n int, matrix [][]int) [][]int {

	size := len(matrix)
	base := make([][][]int, maxlogn)
	base[0] = matrix
	for i := 0; i < maxlogn-1; i++ {
		base[i+1] = matrixmul(base[i], base[i])
	}
	r := i2s(size, size, 0)
	for i := 0; i < size; i++ {
		r[i][i] = 1
	}

	for i := 0; i < maxlogn; i++ {
		if hasbit(n, i) {
			r = matrixmul(r, base[i])
		}
	}
	return r
}

// ==================================================
// point
// ==================================================

type point struct {
	x int
	y int
}

type pointf struct {
	x float64
	y float64
}

func newPoint(x, y int) point {
	return point{x, y}
}

func (p point) isValid(h, w int) bool {
	return 0 <= p.x && p.x < h && 0 <= p.y && p.y < w
}

func (p point) dist(to point) float64 {
	return pointDist(p, to)
}

func pointAdd(a, b point) point {
	return point{x: a.x + b.x, y: a.y + b.y}
}

func pointSub(a, b point) point {
	return point{x: a.x - b.x, y: a.y - b.y}
}

func pointDist(a, b point) float64 {
	return sqrtf((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func pointDistDouble(a, b point) int {
	return (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
}

func pointfDist(a, b pointf) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func pointInnerProduct(a, b point) int {
	return (a.x * b.y) - (b.x * a.y)
}

// ==================================================
// queue
// ==================================================

/*
	q := list.New()
	q.PushBack(val)
	e := q.Front()
	for e != nil {
		t := e.Value.(int)

		// Do something

		e = e.Next()
    }
*/

type IntQueue struct {
	sum   int
	queue []int
	size  int
}

func newIntQueue() *IntQueue {
	return &IntQueue{}
}

func (iq *IntQueue) push(v int) {
	iq.queue = append(iq.queue, v)
	iq.sum += v
	//iq.sum = madd(iq.sum, v)
	iq.size++
}

func (iq *IntQueue) pop() int {
	v := iq.queue[0]
	iq.queue = iq.queue[1:]
	iq.sum -= v
	//iq.sum = madd(iq.sum, -v)
	iq.size--
	return v
}

func (iq *IntQueue) shrink(l int) {
	for {
		if iq.size <= l {
			break
		}
		iq.pop()
	}
}

// ==================================================
// heap
// ==================================================

/*
	ih := newIntHeap(asc)
	ih.Push(v)
	for !ih.IsEmpty() {
		v := ih.Pop()
	}
*/
type IntHeap struct {
	sum int
	pq  *pq
}

func newIntHeap(order sortOrder) *IntHeap {
	ih := &IntHeap{}
	ih.pq = newpq([]compFunc{func(p, q interface{}) int {
		if p.(int) == q.(int) {
			return 0
		}
		if order == asc {
			if p.(int) < q.(int) {
				return -1
			} else {
				return 1
			}
		} else {
			if p.(int) > q.(int) {
				return -1
			} else {
				return 1
			}
		}
	}})
	heap.Init(ih.pq)
	return ih
}
func (ih *IntHeap) Push(x int) {
	ih.sum += x
	heap.Push(ih.pq, x)
}

func (ih *IntHeap) Pop() int {
	v := heap.Pop(ih.pq).(int)
	ih.sum -= v
	return v
}

func (ih *IntHeap) Len() int {
	return ih.pq.Len()
}

func (ih *IntHeap) IsEmpty() bool {
	return ih.pq.Len() == 0
}

func (ih *IntHeap) GetRoot() int {
	return ih.pq.GetRoot().(int)
}

func (ih *IntHeap) GetSum() int {
	return ih.sum
}

/*
	h := &OrgIntHeap{}
	heap.Init(h)

	heap.Push(h, v)
	for !h.IsEmpty() {
		v = heap.Pop(h).(int)
	}
*/
type OrgIntHeap []int

func (h OrgIntHeap) Len() int { return len(h) }

// get from bigger
// func (h OrgIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h OrgIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h OrgIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *OrgIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *OrgIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *OrgIntHeap) IsEmpty() bool {
	return h.Len() == 0
}

// h.Min().(int)
func (h *OrgIntHeap) Min() interface{} {
	return (*h)[0]
}

/*
	type pqst struct {
		x int
		y int
	}

	pq := newpq([]compFunc{func(p, q interface{}) int {
		if p.(pqst).x != q.(pqst).x {
			// get from bigger
			// if p.(pqst).x > q.(pqst).x {
			if p.(pqst).x < q.(pqst).x {
				return -1
			} else {
				return 1
			}
		}
		if p.(pqst).y != q.(pqst).y {
			// get from bigger
			// if p.(pqst).y > q.(pqst).y {
			if p.(pqst).y < q.(pqst).y {
				return -1
			} else {
				return 1
			}
		}
		return 0
	}})
	heap.Init(pq)
	heap.Push(pq, pqst{x: 1, y: 1})
	for !pq.IsEmpty() {
		v := heap.Pop(pq).(pqst)
	}
*/

type pq struct {
	arr   []interface{}
	comps []compFunc
}

type compFunc func(p, q interface{}) int

func newpq(comps []compFunc) *pq {
	return &pq{
		comps: comps,
	}
}

func (pq pq) Len() int {
	return len(pq.arr)
}

func (pq pq) Swap(i, j int) {
	pq.arr[i], pq.arr[j] = pq.arr[j], pq.arr[i]
}

func (pq pq) Less(i, j int) bool {
	for _, comp := range pq.comps {
		result := comp(pq.arr[i], pq.arr[j])
		switch result {
		case -1:
			return true
		case 1:
			return false
		case 0:
			continue
		}
	}
	return true
}

func (pq *pq) Push(x interface{}) {
	pq.arr = append(pq.arr, x)
}

func (pq *pq) Pop() interface{} {
	n := pq.Len()
	item := pq.arr[n-1]
	pq.arr = pq.arr[:n-1]
	return item
}

func (pq *pq) IsEmpty() bool {
	return pq.Len() == 0
}

// pq.GetRoot().(edge)
func (pq *pq) GetRoot() interface{} {
	return pq.arr[0]
}

// ==================================================
// cusum
// ==================================================

type cusum struct {
	l int
	s []int
}

func newcusum(sl []int) *cusum {
	c := &cusum{}
	c.l = len(sl)
	c.s = make([]int, len(sl)+1)
	for i, v := range sl {
		c.s[i+1] = c.s[i] + v
	}
	return c
}

// get sum f <= i && i <= t
func (c *cusum) getRange(f, t int) int {
	if f > t || f >= c.l {
		return 0
	}
	return c.s[t+1] - c.s[f]
}

// get sum 0 to i
func (c *cusum) get(i int) int {
	return c.s[i+1]
}

func (c *cusum) upperBound(i int) int {
	return upperBound(i, c.s)
}

func (c *cusum) lowerBound(i int) int {
	return lowerBound(i, c.s)
}

/*
	mp := make([][]int, n)
	for i := 0; i < k; i++ {
		mp[i] = make([]int, m)
	}
	cusum2d := newcusum2d(sl)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			t:=cusum2d.get(0, 0, i, j)
		}
	}
*/

type cusum2d struct {
	s [][]int
}

func newcusum2d(sl [][]int) *cusum2d {
	c := &cusum2d{}
	n := len(sl)
	m := len(sl[0])
	c.s = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		c.s[i] = make([]int, m+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c.s[i+1][j+1] = c.s[i+1][j] + c.s[i][j+1] - c.s[i][j]
			c.s[i+1][j+1] += sl[i][j]
		}
	}
	return c
}

// x1 <= x <= x2, y1 <= y <= y2
func (c *cusum2d) get(x1, y1, x2, y2 int) int {
	x2++
	y2++
	return c.s[x2][y2] + c.s[x1][y1] - c.s[x1][y2] - c.s[x2][y1]
}

// ==================================================
// union find
// ==================================================

type unionFind struct {
	par     []int
	weights []int
}

func newUnionFind(n int) *unionFind {
	u := &unionFind{
		par:     make([]int, n),
		weights: make([]int, n),
	}
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

func (u *unionFind) root(x int) int {
	if u.par[x] < 0 {
		return x
	}
	px := u.par[x]
	u.par[x] = u.root(px)
	u.weights[x] += u.weights[px]
	return u.par[x]
}

func (u *unionFind) unite(x, y int, arg ...int) {

	w := 0
	if len(arg) == 1 {
		w = arg[0]
	}
	w += u.weight(x)
	w -= u.weight(y)
	x = u.root(x)
	y = u.root(y)
	if x == y {
		return
	}
	if u.size(x) < u.size(y) {
		x, y = y, x
		w = -w
	}
	u.par[x] += u.par[y]
	u.par[y] = x
	u.weights[y] = w
}

func (u *unionFind) issame(x, y int) bool {
	if u.root(x) == u.root(y) {
		return true
	}
	return false
}

func (u *unionFind) size(x int) int {
	return -u.par[u.root(x)]
}

func (u *unionFind) weight(x int) int {
	u.root(x)
	return u.weights[x]
}

func (u *unionFind) diff(x, y int) int {
	return u.weight(y) - u.weight(x)
}

// ==================================================
// bit
// ==================================================

type bit struct {
	n int
	b []int
}

func newbit(n int) *bit {
	return &bit{
		n: n + 1,
		b: make([]int, n+1),
	}
}

func (b *bit) culc(i, j int) int {
	return i + j
	//return madd(i, j)
}

func (b *bit) add(i, x int) {
	for i++; i < b.n && i > 0; i += i & -i {
		b.b[i] = b.culc(b.b[i], x)
	}
}

func (b *bit) sum(i int) int {
	ret := 0
	for i++; i > 0; i -= i & -i {
		ret = b.culc(ret, b.b[i])
	}
	return ret
}

// l <= x < r
func (b *bit) rangesum(l, r int) int {
	return b.culc(b.sum(r-1), -b.sum(l-1))
}

func (b *bit) lowerBound(x int) int {
	idx, k := 0, 1
	for k < b.n {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k < b.n && b.b[idx+k] < x {
			x -= b.b[idx+k]
			idx += k
		}
	}
	return idx
}

// ==================================================
// segment tree
// ==================================================

type streeculctype int

const (
	stadd streeculctype = iota
	stmadd
	stset
	stmin
	stmax
)

/*
s := newstree(n,stmin|stmax|stsum|stmsum)
s.set(i,x)
s.add(i,x)
result1 := s.query(l,r)
result2 := s.findrightest(l,r,x)
result3 := s.findlefttest(l,r,x)
*/
type stree struct {
	n   int
	b   []int
	def int
	cmp func(i, j int) int
}

func newstree(n int, minmax streeculctype) *stree {
	tn := 1
	for tn < n {
		tn *= 2
	}
	s := &stree{
		n: tn,
		b: make([]int, 2*tn-1),
	}
	switch minmax {
	case stmin:
		s.def = inf
		for i := 0; i < 2*tn-1; i++ {
			s.b[i] = s.def
		}
		s.cmp = func(i, j int) int {
			return min(i, j)
		}
	case stmax:
		s.cmp = func(i, j int) int {
			return max(i, j)
		}
	case stadd:
		s.cmp = func(i, j int) int {
			if i == s.def {
				return j
			}
			if j == s.def {
				return i
			}
			return i + j
		}
	case stmadd:
		s.cmp = func(i, j int) int {
			if i == s.def {
				return j
			}
			if j == s.def {
				return i
			}
			return madd(i, j)
		}
	}
	return s
}

func (s *stree) add(i, x int) {
	i += s.n - 1
	s.b[i] += x

	for i > 0 {
		i = (i - 1) / 2
		s.b[i] = s.cmp(s.b[i*2+1], s.b[i*2+2])
	}
}

func (s *stree) set(i, x int) {
	i += s.n - 1
	s.b[i] = x

	for i > 0 {
		i = (i - 1) / 2
		s.b[i] = s.cmp(s.b[i*2+1], s.b[i*2+2])
	}
}

func (s *stree) query(a, b int) int {
	return s.querysub(a, b, 0, 0, s.n)
}

func (s *stree) querysub(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return s.def
	}
	if a <= l && r <= b {
		return s.b[k]
	}
	return s.cmp(
		s.querysub(a, b, k*2+1, l, (l+r)/2),
		s.querysub(a, b, k*2+2, (l+r)/2, r),
	)
}

func (s *stree) findrightest(a, b, x int) int {
	return s.findrightestsub(a, b, x, 0, 0, s.n)
}

func (s *stree) findrightestsub(a, b, x, k, l, r int) int {
	if s.b[k] > x || r <= a || b <= l {
		return a - 1
	} else if k >= s.n-1 {
		return k - s.n + 1
	}
	vr := s.findrightestsub(a, b, x, 2*k+2, (l+r)/2, r)
	if vr != a-1 {
		return vr
	}
	return s.findrightestsub(a, b, x, 2*k+1, l, (l+r)/2)
}

func (s *stree) findleftest(a, b, x int) int {
	return s.findleftestsub(a, b, x, 0, 0, s.n)
}

func (s *stree) findleftestsub(a, b, x, k, l, r int) int {
	if s.b[k] > x || r <= a || b <= l {
		return b
	} else if k >= s.n-1 {
		return k - s.n + 1
	}
	vl := s.findleftestsub(a, b, x, 2*k+1, l, (l+r)/2)
	if vl != b {
		return vl
	}
	return s.findleftestsub(a, b, x, 2*k+2, (l+r)/2, r)
}

func (s *stree) debug() {
	l := []string{}
	t := 2
	out("data")
	for i := 0; i < 2*s.n-1; i++ {
		if i+1 == t {
			t *= 2
			out(strings.Join(l, " "))
			l = []string{}
		}
		if s.b[i] == inf {
			l = append(l, "∞")
		} else {
			l = append(l, strconv.Itoa(s.b[i]))
		}
	}
	out(strings.Join(l, " "))
}

/*
type segstruct struct {
	v    int
	size int
}
*/

type segstruct int

type segfstruct int

type lazysegtree struct {
	n           int
	size        int
	log         int
	d           []segstruct
	lz          []segfstruct
	op          func(segstruct, segstruct) segstruct
	e           func() segstruct
	mapping     func(segfstruct, segstruct) segstruct
	composition func(segfstruct, segfstruct) segfstruct
	id          func() segfstruct
}

/*
	// 区間加算・区間和取得
	op := func(a segstruct, b segstruct) segstruct {
		return segstruct{a.v + b.v, a.size + b.size}
	}
	e := func() segstruct {
		return segstruct{0, 0}
	}
	id := func() segfstruct {
		return segfstruct(inf)
	}
	mapping := func(f segfstruct, x segstruct) segstruct {
		if f == id() {
			return x
		}
		return segstruct{x.v + int(f) * x.size, x.size}
	}
	compostion := func(f segfstruct, g segfstruct) segfstruct {
		if f == id() {
			return g
		}
		if g == id() {
			return f
		}
		return segfstruct(int(f) + int(g))
	}

	// 区間変更・区間最小値取得
	op := func(a segstruct, b segstruct) segstruct {
		return segstruct(min(int(a), int(b)))
	}
	e := func() segstruct {
		return segstruct(inf)
	}
	id := func() segfstruct {
		return segfstruct(inf)
	}
	mapping := func(f segfstruct, x segstruct) segstruct {
		if f == id() {
			return x
		}
		return segstruct(int(f))
	}
	compostion := func(f segfstruct, g segfstruct) segfstruct {
		if i == id() {
			return g
		}
		return f
	}

	lst := newlazysegtree(
		n,
		base,
		op,
		e,
		mapping,
		compostion,
		id,
	)
	lst.applyrange(f, t, segfstruct(v))
	iv := int(lst.get(i))
	rv := int(lst.prod(l,r)
	av := int(t.allprod()))
*/

func newlazysegtree(
	n int,
	v []segstruct,
	op func(segstruct, segstruct) segstruct,
	e func() segstruct,
	mapping func(segfstruct, segstruct) segstruct,
	composition func(segfstruct, segfstruct) segfstruct,
	id func() segfstruct,
) *lazysegtree {

	l := &lazysegtree{
		n:           n,
		op:          op,
		e:           e,
		mapping:     mapping,
		composition: composition,
		id:          id,
	}
	l.size = pow2(bitlen(n))
	l.log = bitlen(n)
	l.d = make([]segstruct, l.size*2)
	for i := range l.d {
		l.d[i] = e()
	}
	l.lz = make([]segfstruct, l.size)
	for i := range l.lz {
		l.lz[i] = id()
	}
	if len(v) > 0 {
		if len(v) != n {
			panic("invalid v value")
		}
		for i := 0; i < l.n; i++ {
			l.d[l.size+i] = v[i]
		}
		for i := l.size - 1; i >= 1; i-- {
			l.update(i)
		}
	}
	return l

}

func (l *lazysegtree) set(p int, x segstruct) {
	if p < 0 || p > l.n {
		panic(fmt.Sprintf("invalid p value n=%v p=%v", l.n, p))
	}
	p += l.size
	for i := l.log; i >= 1; i-- {
		l.push(p >> i)
	}
	l.d[p] = x
	for i := 1; i <= l.log; i++ {
		l.update(p >> i)
	}
}

func (l *lazysegtree) get(p int) segstruct {
	if p < 0 || p > l.n {
		panic(fmt.Sprintf("invalid p value n=%v p=%v", l.n, p))
	}
	p += l.size
	for i := l.log; i >= 1; i-- {
		l.push(p >> i)
	}
	return l.d[p]
}

func (l *lazysegtree) prod(le, ri int) segstruct {
	if le < 0 || le > l.n {
		panic(fmt.Sprintf("invalid le value n=%v ri=%v", l.n, le))
	}
	if ri < 0 || ri > l.n {
		panic(fmt.Sprintf("invalid ri value n=%v ri=%v", l.n, ri))
	}
	if ri < le {
		panic(fmt.Sprintf("invalid le value le=%v ri=%v", le, ri))
	}
	if le == ri {
		return l.e()
	}

	le += l.size
	ri += l.size

	for i := l.log; i >= 1; i-- {
		if ((le >> i) << i) != le {
			l.push(le >> i)
		}
		if ((ri >> i) << i) != ri {
			l.push((ri - 1) >> i)
		}
	}

	sml := l.e()
	smr := l.e()
	for {
		if le >= ri {
			break
		}
		if le&1 == 1 {
			sml = l.op(sml, l.d[le])
			le++
		}
		if ri&1 == 1 {
			ri--
			smr = l.op(l.d[ri], smr)
		}
		le >>= 1
		ri >>= 1

	}

	return l.op(sml, smr)
}

func (l *lazysegtree) allprod() segstruct {
	return l.d[1]
}

func (l *lazysegtree) apply(p int, f segfstruct) {
	if p < 0 || p > l.n {
		panic(fmt.Sprintf("invalid p value n=%v p=%v", l.n, p))
	}
	p += l.size
	for i := l.log; i >= 1; i-- {
		l.push(p >> i)
	}
	l.d[p] = l.mapping(f, l.d[p])
	for i := 1; i <= l.log; i++ {
		l.update(p >> i)
	}
}

func (l *lazysegtree) applyrange(le, ri int, f segfstruct) {
	if le < 0 || le > l.n {
		panic(fmt.Sprintf("invalid le value n=%v ri=%v", l.n, le))
	}
	if ri < 0 || ri > l.n {
		panic(fmt.Sprintf("invalid ri value n=%v ri=%v", l.n, ri))
	}
	if ri < le {
		panic(fmt.Sprintf("invalid le value le=%v ri=%v", le, ri))
	}

	if le == ri {
		return
	}

	le += l.size
	ri += l.size

	for i := l.log; i >= 1; i-- {
		if ((le >> i) << i) != le {
			l.push(le >> i)
		}
		if ((ri >> i) << i) != ri {
			l.push((ri - 1) >> i)
		}
	}

	{
		le2 := le
		ri2 := ri
		for {
			if le >= ri {
				break
			}
			if le&1 == 1 {
				l.allApply(le, f)
				le++
			}
			if ri&1 == 1 {
				ri--
				l.allApply(ri, f)
			}
			le >>= 1
			ri >>= 1
		}
		le = le2
		ri = ri2
	}

	for i := 1; i <= l.log; i++ {
		if ((le >> i) << i) != le {
			l.update(le >> i)
		}
		if ((ri >> i) << i) != ri {
			l.update((ri - 1) >> i)
		}
	}
}

func (l *lazysegtree) maxright(le int, g func(segstruct) bool) int {

	if le < 0 || le > l.n {
		panic(fmt.Sprintf("invalid le value n=%v ri=%v", l.n, le))
	}
	if !g(l.e()) {
		panic("invalid g func")
	}
	if le == l.n {
		return l.n
	}
	le += l.size
	for i := l.log; i >= 1; i-- {
		l.push(le >> i)
	}
	sm := l.e()
	for {
		for {
			if le%2 != 0 {
				break
			}
			le >>= 1
		}
		if !g(l.op(sm, l.d[le])) {
			for {

				if le >= l.size {
					break
				}
				l.push(le)
				le = (2 * le)
				if g(l.op(sm, l.d[le])) {
					sm = l.op(sm, l.d[le])
					le++
				}
			}
			return le - l.size
		}
		sm = l.op(sm, l.d[le])
		le++
		if (le & -le) == le {
			break
		}
	}
	return l.n
}

func (l *lazysegtree) maxleft(ri int, g func(segstruct) bool) int {
	if ri < 0 || ri > l.n {
		panic("invalid ri value")
	}
	if !g(l.e()) {
		panic("invalid g func")
	}

	if ri == 0 {
		return 0
	}
	ri += l.size
	for i := l.log; i >= 1; i-- {
		l.push((ri - 1) >> i)
	}
	sm := l.e()
	for {
		ri--
		for {
			if ri > 1 && (ri%2 == 1) {
			} else {
				break
			}
			ri >>= 1
		}
		if !g(l.op(l.d[ri], sm)) {
			for {
				if ri >= l.size {
					break
				}
				l.push(ri)
				ri = (2*ri + 1)
				if g(l.op(l.d[ri], sm)) {
					sm = l.op(l.d[ri], sm)
					ri--
				}
			}
			return ri + 1 - l.size
		}
		sm = l.op(l.d[ri], sm)
		if (ri & -ri) == ri {
			break
		}
	}
	return 0
}

func (l *lazysegtree) update(k int) {
	l.d[k] = l.op(l.d[2*k], l.d[2*k+1])
}

func (l *lazysegtree) allApply(k int, f segfstruct) {
	l.d[k] = l.mapping(f, l.d[k])
	if k < l.size {
		l.lz[k] = l.composition(f, l.lz[k])
	}
}

func (l *lazysegtree) push(k int) {
	l.allApply(2*k, l.lz[k])
	l.allApply(2*k+1, l.lz[k])
	l.lz[k] = l.id()
}

// ==================================================
// tree
// ==================================================

type tree struct {
	size       int
	root       int
	edges      [][]edge
	parentsize int
	parent     [][]int
	depth      []int
	orderidx   int
	order      []int
}

/*
	n := ni()
	edges := make([][]edge, n)
	for i := 0; i < n-1; i++ {
		s, t := ni2()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t})
		edges[t] = append(edges[t], edge{to: s})
	}
	tree := newtree(n, 0, edges)
	tree.init()
*/
func newtree(size int, root int, edges [][]edge) *tree {
	parentsize := int(math.Log2(float64(size))) + 1
	parent := make([][]int, parentsize)
	for i := 0; i < parentsize; i++ {
		parent[i] = make([]int, size)
	}
	depth := make([]int, size)
	order := make([]int, size)
	return &tree{
		size:       size,
		root:       root,
		edges:      edges,
		parentsize: parentsize,
		parent:     parent,
		depth:      depth,
		order:      order,
	}
}

func (t *tree) init() {
	t.dfs(t.root, -1, 0)
	for i := 0; i+1 < t.parentsize; i++ {
		for j := 0; j < t.size; j++ {
			if t.parent[i][j] < 0 {
				t.parent[i+1][j] = -1
			} else {
				t.parent[i+1][j] = t.parent[i][t.parent[i][j]]
			}
		}
	}
}

func (t *tree) dfs(v, p, d int) {
	t.order[v] = t.orderidx
	t.orderidx++
	t.parent[0][v] = p
	t.depth[v] = d
	for _, nv := range t.edges[v] {
		if nv.to != p {
			t.dfs(nv.to, v, d+1)
		}
	}
}

func (t *tree) lca(u, v int) int {
	if t.depth[u] > t.depth[v] {
		u, v = v, u
	}
	for i := 0; i < t.parentsize; i++ {
		if (t.depth[v]-t.depth[u])>>i&1 == 1 {
			v = t.parent[i][v]
		}
	}
	if u == v {
		return u
	}
	for i := t.parentsize - 1; i >= 0; i-- {
		if t.parent[i][u] != t.parent[i][v] {
			u = t.parent[i][u]
			v = t.parent[i][v]
		}
	}
	return t.parent[0][u]
}

func (t *tree) dist(u, v int) int {
	return t.depth[u] + t.depth[v] - t.depth[t.lca(u, v)]*2
}

func (t *tree) auxiliarytree(sl []int) []int {
	sort.Slice(sl, func(i, j int) bool { return t.order[sl[i]] < t.order[sl[j]] })
	return sl
}

// ==================================================
// graph
// ==================================================

type edge struct {
	from int
	to   int
	cost int
	rev  int
}

func setDualEdge(edges [][]edge, s, t, c int) {
	edges[s] = append(edges[s], edge{to: t, cost: c, rev: len(edges[t])})
	edges[t] = append(edges[t], edge{to: s, cost: 0, rev: len(edges[s]) - 1})
}

func reverseEdgeCost(edges [][]edge, from, i int) {
	redge := edges[from][i]
	t := edges[redge.to][redge.rev].cost
	edges[redge.to][redge.rev].cost = redge.cost
	edges[redge.from][i].cost = t
}

func eraseEdgeCost(edges [][]edge, from, i int) {
	redge := edges[from][i]
	edges[redge.to][redge.rev].cost = 0
	edges[redge.from][i].cost = 0
}

type state struct {
	score int
	node  int
}

type graph struct {
	size         int
	edges        [][]edge
	starts       []state
	comps        []compFunc
	defaultScore int
	level        []int
	iter         []int
}

func newgraph(size int, edges [][]edge) *graph {
	graph := &graph{
		size:  size,
		edges: edges,
	}

	graph.defaultScore = inf
	graph.comps = []compFunc{
		func(p, q interface{}) int {
			if p.(state).score < q.(state).score {
				return -1
			} else if p.(state).score == q.(state).score {
				return 0
			}
			return 1
		},
	}
	return graph
}

/*
	v, e := ni2()
	edges := make([][]edge, v)
	deg := make([]int, v)
	for i := 0; i < e; i++ {
		s, t, c := ni3()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t, cost: c})
		deg[t]++
	}
	graph := newgraph(v, edges)
	isdag, r := graph.topologicalSort(deg)
*/
func (g *graph) topologicalSort(deg []int) (bool, []int) {

	r := []int{}
	q := list.New()
	for i := 0; i < g.size; i++ {
		if deg[i] == 0 {
			q.PushBack(i)
		}
	}
	e := q.Front()
	for e != nil {
		t := e.Value.(int)
		r = append(r, t)
		for _, edge := range g.edges[t] {
			deg[edge.to]--
			if deg[edge.to] == 0 {
				q.PushBack(edge.to)
			}
		}

		e = e.Next()
	}
	for _, v := range deg {
		if v != 0 {
			return false, nil
		}
	}
	return true, r
}

/*
	v, e := ni2()
	edges := make([][]edge, v)
	edgers := make([][]edge, v)

	for i := 0; i < e; i++ {
		s, t := ni2()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t})
		edgers[t] = append(edgers[t], edge{to: s})
	}

	scc := getScc(v, edges, edgers)
*/
func getScc(v int, edges, edgers [][]edge) [][]int {
	used := make([]bool, v)
	scc := [][]int{}
	vs := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		used[i] = true
		for _, v := range edges[i] {
			if used[v.to] == false {
				dfs(v.to)
			}
		}
		vs = append(vs, i)
	}

	var rdfs func(i, k int)
	rdfs = func(i, k int) {
		used[i] = true
		scc[k] = append(scc[k], i)
		for _, v := range edgers[i] {
			if used[v.to] == false {
				rdfs(v.to, k)
			}
		}
	}

	for i := 0; i < v; i++ {
		if used[i] == false {
			dfs(i)
		}
	}
	used = make([]bool, v)
	k := 0
	for i := v - 1; i >= 0; i-- {
		if used[vs[i]] == false {
			scc = append(scc, []int{})
			rdfs(vs[i], k)
			k++
		}
	}
	return scc
}

/*
	v, e := ni2()
	edges := make([][]edge, v)

	for i := 0; i < e; i++ {
		s, t, c := ni3()
		s--
		t--
		edges[s] = append(edges[s], edge{to: t, cost: c})
		edges[t] = append(edges[t], edge{to: s, cost: c})
	}

	graph := newgraph(v, edges)
	dist := graph.dijkstra(0)
*/

func (g *graph) dijkstra(start int) []int {

	g.starts = []state{{node: start}}

	score := make([]int, g.size)
	for i := 0; i < g.size; i++ {
		score[i] = g.defaultScore
	}
	que := newpq(g.comps)
	for _, start := range g.starts {
		score[start.node] = start.score
		heap.Push(que, start)
	}

	for !que.IsEmpty() {
		st := heap.Pop(que).(state)
		if st.score > score[st.node] {
			continue
		}
		for _, edge := range g.edges[st.node] {
			newScore := st.score + edge.cost
			if score[edge.to] > newScore {
				score[edge.to] = newScore
				heap.Push(que, state{score: newScore, node: edge.to})
			}
		}
	}
	return score
}

func (g *graph) floydWarshall() ([][]int, bool) {

	score := make([][]int, g.size)
	for i := 0; i < g.size; i++ {
		score[i] = make([]int, g.size)
		for j := 0; j < g.size; j++ {
			if i == j {
				score[i][j] = 0
			} else {
				score[i][j] = g.defaultScore
			}
		}
		for _, edge := range g.edges[i] {
			score[i][edge.to] = edge.cost
		}
	}
	for k := 0; k < g.size; k++ {
		for i := 0; i < g.size; i++ {
			for j := 0; j < g.size; j++ {
				if score[i][k] == g.defaultScore || score[k][j] == g.defaultScore {
					continue
				}
				score[i][j] = min(score[i][j], score[i][k]+score[k][j])
			}
		}
	}

	for k := 0; k < g.size; k++ {
		if score[k][k] < 0 {
			return nil, true
		}
	}

	return score, false
}

/*
	v, e := ni2()
	edges := make([][]edge, 1)
	edges[0] = make([]edge, e)

	for i := 0; i < e; i++ {
		s, t, d := ni3()
		edges[0][i] = edge{from: s, to: t, cost: d}
	}

	graph := newgraph(v, edges)

	o = graph.kruskal()
*/
func (g *graph) kruskal() int {

	sort.Slice(g.edges[0], func(i, j int) bool { return g.edges[0][i].cost < g.edges[0][j].cost })

	e := len(g.edges[0])

	uf := newUnionFind(g.size)
	r := 0
	for i := 0; i < e; i++ {
		edge := g.edges[0][i]
		if uf.issame(edge.from, edge.to) {
			continue
		}
		r += edge.cost
		uf.unite(edge.from, edge.to)
	}

	return r
}

/*
	v, e := ni2()
	edges := make([][]edge, v)
	for i := 0; i < e; i++ {
		s, t, c := ni3()
		s--
		t--
		setDualEdge(edges, s, t, c)
	}
	graph := newgraph(v, edges)
	o = graph.dinic()
*/
func (g *graph) dinic() int {
	f := 0
	for {
		g.dinicbfs(0)
		if g.level[g.size-1] < 0 {
			break
		}
		g.iter = make([]int, g.size)
		for {
			t := g.dinicdfs(0, g.size-1, inf)
			if t <= 0 {
				break
			}
			f += t
		}
	}
	return f
}

func (g *graph) dinicbfs(s int) {
	g.level = make([]int, g.size)
	for i := 0; i < g.size; i++ {
		g.level[i] = -1
	}
	g.level[s] = 0

	q := []int{}
	q = append(q, s)
	ti := 0
	for {
		if ti >= len(q) {
			break
		}
		t := q[ti]

		for _, e := range g.edges[t] {
			if e.cost > 0 && g.level[e.to] < 0 {
				g.level[e.to] = g.level[t] + 1
				q = append(q, e.to)
			}
		}

		ti++
	}
}

func (g *graph) dinicdfs(v, t, f int) int {
	if v == t {
		return f
	}
	for i := g.iter[v]; i < len(g.edges[v]); i++ {
		e := g.edges[v][i]
		g.iter[v] = i

		if e.cost > 0 && g.level[v] < g.level[e.to] {
			d := g.dinicdfs(e.to, t, min(f, e.cost))
			if d > 0 {
				g.edges[v][i].cost -= d
				g.edges[e.to][e.rev].cost += d
				return d
			}
		}
	}
	return 0
}

// ==================================================
// fft
// ==================================================

func convolution(a, b []int) []int {
	n1, n2 := len(a), len(b)
	n := n1 + n2 - 1
	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	MOD1 := 754974721
	MOD2 := 167772161
	MOD3 := 469762049
	M2M3 := MOD2 * MOD3
	M1M3 := MOD1 * MOD3
	M1M2 := MOD1 * MOD2
	M1M2M3 := MOD1 * MOD2 * MOD3

	i1 := minv(M2M3, MOD1)
	i2 := minv(M1M3, MOD2)
	i3 := minv(M1M2, MOD3)

	c1 := convolutionMod(a, b, MOD1)
	c2 := convolutionMod(a, b, MOD2)
	c3 := convolutionMod(a, b, MOD3)

	c := make([]int, n)
	offset := []int{0, 0, M1M2M3, 2 * M1M2M3, 3 * M1M2M3}

	for i := 0; i < n; i++ {
		x := 0
		x += c1[i] * i1 % MOD1 * M2M3
		x += c2[i] * i2 % MOD2 * M1M3
		x += c3[i] * i3 % MOD3 * M1M2
		diff := c1[i] - x%MOD1
		if diff < 0 {
			diff += MOD1
		}
		x -= offset[diff%5]
		c[i] = x
	}

	return c
}

func convolutionMod(a, b []int, mod int) []int {
	n1, n2 := len(a), len(b)
	n := n1 + n2 - 1
	if n1 == 0 || n2 == 0 {
		return []int{}
	}

	z := 1 << ceilPow2(n)
	aa, bb := make([]int, z), make([]int, z)
	copy(aa, a)
	copy(bb, b)
	a, b = aa, bb

	butterfly(a, mod)
	butterfly(b, mod)
	for i := 0; i < z; i++ {
		a[i] = a[i] * b[i] % mod
	}
	butterflyInv(a, mod)
	a = a[:n]
	iz := minv(z, mod)
	for i := 0; i < n; i++ {
		a[i] = a[i] * iz % mod
		if a[i] < 0 {
			a[i] += mod
		}
	}

	return a
}

func primitiveRoot(m int) int {
	if m == 2 {
		return 1
	}
	if m == 167772161 || m == 469762049 || m == 998244353 {
		return 3
	}
	if m == 754974721 {
		return 11
	}
	divs := make([]int, 20)
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for x%2 == 0 {
		x /= 2
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := 2; ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			if mpow(g, (m-1)/divs[i], m) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
}

func ceilPow2(n int) int {
	x := 0
	for 1<<x < n {
		x++
	}
	return x
}

func bsf(n int) int {
	x := 0
	for n&(1<<x) == 0 {
		x++
	}
	return x
}

func butterfly(a []int, M int) {
	g := primitiveRoot(M)
	n := len(a)
	h := ceilPow2(n)

	se := make([]int, 30)
	es, ies := make([]int, 30), make([]int, 30)
	cnt2 := bsf(M - 1)
	e := mpow(g, (M-1)>>cnt2, M)
	ie := minv(e, M)
	for i := cnt2; i >= 2; i-- {
		es[i-2] = e
		ies[i-2] = ie
		e = e * e % M
		ie = ie * ie % M
	}
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		se[i] = es[i] * now % M
		now = now * ies[i] % M
	}
	for ph := 1; ph <= h; ph++ {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		now := 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := a[i+offset+p] * now % M
				a[i+offset] = (l + r) % M
				a[i+offset+p] = (M + l - r) % M
			}
			now = now * se[bsf(^s)] % M
		}
	}
}

func butterflyInv(a []int, M int) {
	g := primitiveRoot(M)
	n := len(a)
	h := ceilPow2(n)

	sie := make([]int, 30)
	es, ies := make([]int, 30), make([]int, 30)
	cnt2 := bsf(M - 1)
	e := mpow(g, (M-1)>>cnt2, M)
	ie := minv(e, M)
	for i := cnt2; i >= 2; i-- {
		es[i-2] = e
		ies[i-2] = ie
		e = e * e % M
		ie = ie * ie % M
	}
	now := 1
	for i := 0; i <= cnt2-2; i++ {
		sie[i] = ies[i] * now % M
		now = now * es[i] % M
	}
	for ph := h; ph >= 1; ph-- {
		w := 1 << (ph - 1)
		p := 1 << (h - ph)
		inow := 1
		for s := 0; s < w; s++ {
			offset := s << (h - ph + 1)
			for i := 0; i < p; i++ {
				l := a[i+offset]
				r := a[i+offset+p]
				a[i+offset] = (l + r) % M
				a[i+offset+p] = (M + l - r) * inow % M
			}
			inow = inow * sie[bsf(^s)] % M
		}
	}
}
