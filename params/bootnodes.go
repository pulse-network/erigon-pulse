// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	libcommon "github.com/ledgerwatch/erigon-lib/common"

	"github.com/ledgerwatch/erigon/params/networkname"
)

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",   // bootnode-aws-ap-southeast-1-001
	"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",     // bootnode-aws-us-east-1-001
	"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",   // bootnode-azure-australiaeast-001
	"enode://103858bdb88756c71f15e9b5e09b56dc1be52f0a5021d46301dbbfb7e130029cc9d0d6f73f693bc29b665770fff7da4d34f3c6379fe12721b5d7a0bcb5ca1fc1@191.234.162.198:30303", // bootnode-azure-brazilsouth-001
	"enode://715171f50508aba88aecd1250af392a45a330af91d7b90701c436b618c86aaa1589c9184561907bebbb56439b8f8787bc01f49a7c77276c58c1b09822d75e8e8@52.231.165.108:30303",  // bootnode-azure-koreasouth-001
	"enode://5d6d7cd20d6da4bb83a1d28cadb5d409b64edf314c0335df658c1a54e32c7c4a7ab7823d57c39b6a757556e68ff1df17c748b698544a55cb488b52479a92b60f@104.42.217.25:30303",   // bootnode-azure-westus-001
	"enode://2b252ab6a1d0f971d9722cb839a42cb81db019ba44c08754628ab4a823487071b5695317c8ccd085219c3a03af063495b2f1da8d18218da2d6a82981b45e6ffc@65.108.70.101:30303",   // bootnode-hetzner-hel
	"enode://4aeb4ab6c14b23e2c4cfdce879c04b0748a20d8e9b59e25ded2a08143e265c6c25936e74cbc8e641e3312ca288673d91f2f93f8e277de3cfa444ecdaaf982052@157.90.35.166:30303",   // bootnode-hetzner-fsn
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	"enode://9246d00bc8fd1742e5ad2428b80fc4dc45d786283e05ef6edbd9002cbc335d40998444732fbe921cb88e1d2c73d1b1de53bae6a2237996e9bfe14f871baf7066@18.168.182.86:30303",
	"enode://ec66ddcf1a974950bd4c782789a7e04f8aa7110a72569b6e65fcd51e937e74eed303b1ea734e4d19cfaec9fbff9b6ee65bf31dcb50ba79acce9dd63a6aca61c7@52.14.151.177:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://d4f764a48ec2a8ecf883735776fdefe0a3949eb0ca476bd7bc8d0954a9defe8fea15ae5da7d40b5d2d59ce9524a99daedadf6da6283fca492cc80b53689fb3b3@46.4.99.122:32109",
	"enode://d2b720352e8216c9efc470091aa91ddafc53e222b32780f505c817ceef69e01d5b0b0797b69db254c586f493872352f5a022b4d8479a00fc92ec55f9ad46a27e@88.99.70.182:30303",
}

var BscBootnodes = []string{
	"enode://1cc4534b14cfe351ab740a1418ab944a234ca2f702915eadb7e558a02010cb7c5a8c295a3b56bcefa7701c07752acd5539cb13df2aab8ae2d98934d712611443@52.71.43.172:30311",
	"enode://28b1d16562dac280dacaaf45d54516b85bc6c994252a9825c5cc4e080d3e53446d05f63ba495ea7d44d6c316b54cd92b245c5c328c37da24605c4a93a0d099c4@34.246.65.14:30311",
	"enode://5a7b996048d1b0a07683a949662c87c09b55247ce774aeee10bb886892e586e3c604564393292e38ef43c023ee9981e1f8b335766ec4f0f256e57f8640b079d5@35.73.137.11:30311",
}
var SepoliaStaticPeers = []string{
	// from https://github.com/ledgerwatch/erigon/issues/6134#issuecomment-1354923418
	"enode://8ae4559db1b1e160be8cc46018d7db123ed6d03fbbfe481da5ec05f71f0aa4d5f4b02ad059127096aa994568706a0d02933984083b87c5e1e3de2b7692444d37@35.161.233.158:46855",
	"enode://d0b3b290422f35ec3e68356f3a4cdf9c661f71a868110670e31441a5021d7abd0440ae8dfb9360aafdd0198f177863361e3a7a7eb5e1a3e26575bf1ac3ef4ab3@162.19.136.65:48264",
	"enode://d64624bda3cdb65d542c90757a4a661cfe9dddf8328bdb1ea97a8d70fad287c360f0101c492d8fd6ab30d79160a3bf148cacfd68f5d2e47eab0b709516419304@51.195.63.10:30040",
	"enode://c7df835939e027325c6bba926220fae5912a33c83d96b3eef8ef445c98083f3191788581c9a0e8f74cadb0b13229b847f5c1ebd315b22bcf11faf6468020eb48@54.163.51.157:30303",
	"enode://da0609bad3afcab9b93175a41a2d621d07aa7ff6c134a00792d4541f0ce8d30d8f3c51bb37a47573508a0bf18865b04066af2a661edf1d3a3d8d133fc1031aa0@88.151.101.14:45192",
	"enode://7a4534d392c59369eae6befa56ac670476d9edc16597cf53c92bbefa6e741b6b0b9e6822cab12afb09123e03ca1131026fbef145adec429fe2e50182dfb650a5@94.130.18.108:31312",
	"enode://db6fa13b63a885440de581ee3fc8df9c6a590326b39fc5ccba7991707ee0cebac306211f7eca5270a350201a3132511f2338481edd81f3dc819c2a1c60419cf2@65.21.89.157:30303",
	"enode://fcf03e9404cace34c60e4eed374ef9a779471014319b3346352fbc2f992a399af6517486e8e65a4ab55f4645fe55420bbea1cddc13a4af4df63b0f731915c6a6@13.125.238.49:46173",
	"enode://8b973816278fdd56966709e4794c7ccce1f256eaa9165a6b013b991a9bdf3886a8f2d23af50ee723a5614a9fe9d197252b803b4455a87ab2468e128f7b06e0ca@172.104.107.145:30303",
	"enode://5a1fb15f826a213d3ef4adb9be47ab58b2240ea05df0d760a244f04762b0847dcb08276b1284f726c22eea30fce0c601cf121b81bac0c151f1b3b4ad00d1482a@34.159.55.147:51262",
	"enode://560928dd14819f88113586726e452b16bbc694ed4144ddadd6290053e7f3fc66bfad13add6889f7d8f37e0c21ccbb6948eb8899c8b30743f4b45a3081f1efed8@34.138.254.5:29888",
	"enode://69a13b575b8c5278431409e9f7db36e7218667ae286bfb65a72dfec9201b2c5bbbe2797a1babbdf17a7bf7ca68fa3fbe1554612637eb1b2425fa975e1bccb54c@35.223.41.3:30303",
	"enode://66158b31eecff939f220b291d2b448edbfe94f1d4c992d9395b5d476e55e54b5abd11d3ee44daf1e18ee27b910ef99cdf6f19775eb4820ebe4f77d7aa948e3b6@51.195.63.10:55198",
	"enode://bf94acbd51170bf075cacb9f149b21ff46354d659ab434a0d40688f776e1e1556bc62be2dc2867ba513844268c0dc8240099a6b60efe1713fbc25da7fdeb6ff1@3.82.105.139:30303",
	"enode://41329e5ceb51cdddbe6a475db00b682505768b71ff8ee37d2d3500ca1b78918f9fad57d6006dd9f79cd418437dbcf87ec2fd58d60710f925cb17da05a51197cf@65.21.34.60:30303",
}

var BscStaticPeers = []string{
	"enode://c688c9a868374ec4fd99f3bf668bff6f1b3c4f7ef6ffa3e908897e4fdc7be99f239d9d3e698db9bc12ae41291b6d32eeaef1d5665662d84b4cf09756e54e4052@3.226.45.227:30311",
	"enode://d61a31410c365e7fcd50e24d56a77d2d9741d4a57b295cc5070189ad90d0ec749d113b4b0432c6d795eb36597efce88d12ca45e645ec51b3a2144e1c1c41b66a@34.204.129.242:30311",
	"enode://c753758ca72c55dd389c270d4f94dbb5ef6141649f0e3dabe515498378dde9c0363deb8614ad2568a8f020eb02de5eea41f9116eed735972be9ffa452a1ff242@54.81.254.21:30311",
	"enode://cfbdc91fe451b4ec95b3b96a379a4b3464ca7d33fe5f0d6783c45dcfaf4431a22a176d2d8aeea6f6309d7c3153e4b307de06aac52b76e39c4146840bc4956bfb@54.236.113.206:30311",
	"enode://bf081097931e2a5e092d6df2e903dab87551d4a6ce3b3da1b3df5c1ee705fc3e3191266e6227fea058ce4e1299cf7092f7dee80df2eadced92fd8263406b0b64@34.237.238.206:30311",
	"enode://1cc4534b14cfe351ab740a1418ab944a234ca2f702915eadb7e558a02010cb7c5a8c295a3b56bcefa7701c07752acd5539cb13df2aab8ae2d98934d712611443@100.24.179.110:30311",
	"enode://c09132e0ac4875c1b0ac2856e032875f174debc76c779b125ec2dc44e94ae42b6db4412b970f15a6e618a48a6ffff6f42230222d6e61e30fd4426a1af82b154b@54.85.57.23:30311",
	"enode://c6ffd994c4ef130f90f8ee2fc08c1b0f02a6e9b12152092bf5a03dd7af9fd33597d4b2e2000a271cc0648d5e55242aeadd6d5061bb2e596372655ba0722cc704@54.147.151.108:30311",
	"enode://786acbdf5a3cf91b99047a0fd8305e11e54d96ea3a72b1527050d3d6f8c9fc0278ff9ef56f3e56b3b70a283d97c309065506ea2fc3eb9b62477fd014a3ec1a96@107.23.90.162:30311",
	"enode://d257b0f8d362161bb4a18b13253b1f6abb273a89d85fed73458f84dbf52d14ce556fd32ac07cbc49cd69d99c03ceddeccab919e10351c9d4e795a4c36404aef7@3.218.112.8:30311",
	"enode://2aadea249d51e076060c97cbc293a2d90052d1786062f137aa8f4afe06996b674ecdbe660d3033b3d46d9cdaf2c3e3db6c031e218edc5a2613caa7f9705774f9@34.193.176.125:30311",
	"enode://0e8b399223f64479d5022b76e8f025214183d0a01b7f044a38a1844f57e9fe89d5da1fa04071758d0c39fe6afba4d5798703b4200a5078664c488129ff8239b7@18.207.56.184:30311",
	"enode://742a02a0d8612e6c4a3669f525a9a045e5c7b0d38ab728ca7086901c5f00b90cab8d796eff64dcfa516cff4e45591f8cb3a1e34277810f84e22ac72311c42cd8@52.21.104.255:30311",
	"enode://9f2246671f19410e722b4f977e0f2997a95a1614fafccc2db15c9dee810bfb10628d3ecc0f6eed5c2cbc1f7fc71b5fb5eb3afcd533cba5f02a6bda673f99ce29@54.204.215.58:30311",
	"enode://bcccac7dfbd21bd1dbbb7bf64ef7af6986520091f00320905e5919d12167fc8c94698099946933747ad2ea24745d35ebc980f9ba1475a839c83574d5ccf318d9@34.196.94.250:30311",
	"enode://20d04257749893d7193b8e3ed619d46384d28b350508bef163b52ee9dc60efc4f562aee00c7fde5cfa83e4e9723b0e90d6422d9031b6069734bd7e24a9ed8e73@107.21.209.99:30311",

	//from https://gist.github.com/roderik/60fb0e376518fabaed6d703096269353
	"enode://9f90d69c5fef1ca0b1417a1423038aa493a7f12d8e3d27e10a5a8fd3da216e485cf6c15f48ee310a14729bc3a4b05038479476c0aa82eed3c5d9d2e64ba3a2b3@52.69.42.169:30311",
	"enode://78ef719ebb2f4fc222aa988a356274dcd3624fb808936ca2ea77388ca229773d4351f795abf505e86db1a30ed1523ded9f9674d916b295bfb98516b78d2844be@13.231.200.147:30311",
	"enode://a8ff9670029785a644fb709ec7cd7e7e2d2b93761872bfe1b011a1ed1c601b23ffa69ead0901b759d780ed65aa81444261905b6964bdf8647bf5b061a4796d2d@54.168.191.244:30311",
	"enode://0f0abad52d6e3099776f70fda913611ad33c9f4b7cafad6595691ea1dd57a37804738be65315fc417d41ab52632c55a5f5f1e5ed3123ed64a312341a8c3f9e3c@52.193.230.222:30311",
	"enode://ecc277f466f35b249b62de8ca567dfe759162ffecc79f40339655537ee58132aec892bc0c4ad3dfb0ba5441bb7a68301c0c09e3f66454110c2c03ccca084c6b5@54.238.240.9:30311",
	"enode://dd3fb5f4da631067d0a9206bb0ac4400d3a076102194257911b632c5aa56f6a3289a855cc0960ad7f2cda3ba5162e0d879448775b07fa73ccd2e4e0477290d9a@54.199.96.72:30311",
	"enode://74481dd5079320755588b5243f82ddec7364ad36108ac77272b8e003194bb3f5e6386fcd5e50a0604db1032ac8cb9b58bb813f8e57125ad84ec6ceec65d29b4b@52.192.99.35:30311",
	"enode://190df80c16509d9d205145495f169a605d1459e270558f9684fcd7376934e43c65a38999d5e49d2ad118f49abfb6ff62068051ce49acc029da7d2be9910fe9fd@13.113.113.139:30311",
	"enode://368fc439d8f86f459822f67d9e8d1984bab32098096dc13d4d361f8a4eaf8362caae3af86e6b31524bda9e46910ac61b075728b14af163eca45413421386b7e2@52.68.165.102:30311",
	"enode://2038dac8d835db7c4c1f9d2647e37e6f5c5dc5474853899adb9b61700e575d237156539a720ff53cdb182ee56ac381698f357c7811f8eadc56858e0d141dcce0@18.182.11.67:30311",
	"enode://fc0bb7f6fc79ad7d867332073218753cb9fe5687764925f8405459a98b30f8e39d4da3a10f87fe06aa10df426c2c24c3907a4d81df4e3c88e890f7de8f8980de@54.65.239.152:30311",
	"enode://3aaaa0e0c7961ef3a9bf05f879f84308ca59651327cf94b64252f67448e582dcd6a6dbe996264367c8aa27fc302736db0283a3516c7406d48f268c5e317b9d49@34.250.1.192:30311",
	"enode://62c516645635f0389b4c851bfc4545720fac0607de74942e4ea7e923f4fa2ac0c438c146e2f0721c8ce06dca4e7f30f5c0136569d9f4b6a827c62b980fd53272@52.215.57.20:30311",
	"enode://5df2f71ae6b2e3bb92f92badbce1f601feabd2d6ce899cf8265c39c38ff446136d74f5bfa089532c7074bb7606a509a54a2ac66397aaaab2363dad3f43c687a8@79.125.103.83:30311",
	"enode://760b5fde9bc14155fa2a87e56cf610701ad6c1adcf44555a7b839baf71f86f11cdadcaf925e50b17c98cc28e20e0df3c3463caad7c6658a76ab68389af639f33@34.243.1.225:30311",
	"enode://57824d2d9b5f39681bee265d56ec98a17fa4af343debdeba18596837f776f7c6370d8a33354e2b1750c41b221778e05c4189b93aca0d4cb1d45d32dc3b2d63f1@34.240.198.163:30311",
	"enode://9b7ff9e2d2154f6de3f53db2123e6f9a6b5b29414d9d5ae8277592b361158c25fcab86e6bfad5ef6554c6d92fb4ca897f7342563e355b80bcdc994f9c268dc2f@34.251.95.115:30311",
	"enode://67ec1f3df346e0aef401175119172e86a20e7ee1442cba4a2074519405cdae3708be3fdcb5e139094408b5d6f6c8e85f89ebb77d04833f7aa251c91344dbd4c9@3.249.178.199:30311",
	"enode://99c8d55d4528330fc494705ea15c2a8be9c25cb638ed561657a642d57e7851e38365d20b6864419e82e593e2b8d22cee23a09e9bb774ec8f15795b077bae7aeb@54.229.26.251:30311",
	"enode://1afc9727301dcd8d2c5aef067031639ae3d3c7a23f8ba6c588a6a1b2c3cbcd738b4ccc53c07d08690ef591b99fd12f00a005f38d820354a91f418ab0939b9072@34.253.216.225:30311",
	"enode://7c7b46ad65325f16768013167a0b2ca3eaa20e5d594011b6202b9c4707f740e2c795e84563b3a8c7986fdfb413ce88726a096f3cac8366ac9ebf073095c20584@34.243.12.13:30311",
	"enode://71ef36f019bbdaa2a7b4676a61d014d0be81958e2c60dd95c66a5e1af10de6f3a62ecf9ad0c26b6c5789b81ac22f774abb4735cd9e259185773ebfd1efded5de@54.170.254.50:30311",
	"enode://627a1cb2c4712cce439026da0c2f599b97628c90c8ccc55526574a944b7455827544130b3003e79399cd79bd73a06a1d6bbd018fcf9ffc5297d3b731aa1b40ab@3.91.73.29:30311",
	"enode://16c7e98f78017dafeaa4129647d1ec66b32ee9be5ec753708820b7363091ceb310f575e7abd9603005e0e34d7b3316c1a4b6c8c42d7f074ed2eb4d073f800a03@3.85.216.212:30311",
	"enode://accbc0a5af0af03e1ec3b5e80544bdceea48011a6928cd82d2c1a9c38b65fd48ec970ba17bd8c0b0ec21a28faec9efe1d1ce55134784b9207146e2f62d8932ba@54.162.32.1:30311",
	"enode://c64c864572dae7ea25225a412c026ced0de66ae429b40c545be8f524a1aeb70b3441710dbfed19e3ba9ef08ce13b00a58daa7a7510924da8e6f4f412d8b45fd5@3.92.160.2:30311",
	"enode://5a838185d4b91eb42cbe3a60bb9f706484d8ec5041fa97b557d10e8ca10a459db0271e06e8b85cad57f1d2c7b05aa4319c0300b2936eefcb2302e10b253cf7d6@23.20.67.34:30311",
	"enode://3438d60bcb628ba33b0adf5e653751436fdc393a869fab136dec5ec6b2ed06d8ea30e4fec061f4f4a67bb01644897dbc3d14db44afc052eb69f102340aff70f9@18.215.252.114:30311",
	"enode://c307b4cddec0aea2188eafddedb0a076b9289402c63217b4c81eb7f34761c7cfaf6b075e93d7357169e226ff1bb4aa3bd71869b4c76cf261e2991005ddb4d4aa@3.81.81.182:30311",
	"enode://80f446f15c3c17b2f8cd7e0f7811f9ba62381abeabc0ce562134d6ac7d400aef212020c439f462d760ca250e8f14b50f215d65e7137d2e3e25d22dc8ff21bda7@54.162.73.225:30311",
	"enode://d69853daf3057cc191514afdf56df4769238fde4f261fab80c6e089480abb9916d61180e783d1cc9e5ae56d30ce6261d9954702dc73c41cd47e4b3961830b2dc@184.73.34.17:30311",
	"enode://ba88d1a8a5e849bec0eb7df9eabf059f8edeae9a9eb1dcf51b7768276d78b10d4ceecf0cde2ef191ced02f66346d96a36ca9da7d73542757d9677af8da3bad3f@54.198.97.197:30311",
	"enode://a232f92d1e76447b93306ece2f6a55ac70ca4633fae0938d71a100757eaf8526e6bbf720aa70cba1e6d186be17291ad1ee851a35596ec6caa2fdf135ce4b6b68@107.20.124.16:30311",
	"enode://2d55e48679442a9e3ef2a3edf2854dcb289f8162d57dbda1e82e7576b0708e0670befaa7255f5c9fa8389443a7e7b4ff762c9e7fd33ddf9f21ec9562f03e8945@18.212.135.123:30311",
	"enode://f7dc512940ca4a8f6858632abbdfc59cea6c4ed7a8da41ddfc4e4dac74e2664e74355fd7c688b285a22295e0053a800f759c9123ec741285a5bd602f89720cea@54.198.51.232:30311",
	"enode://9df97e190f0b82ba7891e0ed556f11f4c1a172c26b2e823e52cfe5722b3df3f1819d2acb87ed0bfeb21fe3aee4ef1ffb8c9227fa7fdf744bfd4f47caad461edf@54.81.89.198:30311",
	"enode://55ef168aab6fe8e14bb536a6c5d1d7c9330bad75a597f7e2cfc6233e6e2a6e58fb341ee17917f494f88d7393f191a042be40af7ca43fb6150b0f89280211db28@54.75.119.68:30311",
	"enode://ae74385270d4afeb953561603fcedc4a0e755a241ffdea31c3f751dc8be5bf29c03bf46e3051d1c8d997c45479a92632020c9a84b96dcb63b2259ec09b4fde38@18.180.28.21:30311",
	"enode://d257b0f8d362161bb4a18b13253b1f6abb273a89d85fed73458f84dbf52d14ce556fd32ac07cbc49cd69d99c03ceddeccab919e10351c9d4e795a4c36404aef7@3.218.112.8:30311",
	"enode://faf14b1c38422f9576760b4d86b62ee6d2b436d1cac9965adf3f5af23dece3a9829ad3f53015cda992c6bbdc44816ccc96043fa342b839308681909cbd0ea9b2@52.71.45.253:30311",
	"enode://f420209bac5324326c116d38d83edfa2256c4101a27cd3e7f9b8287dc8526900f4137e915df6806986b28bc79b1e66679b544a1c515a95ede86f4d809bd65dab@54.178.62.117:30311",
	"enode://e421c47edc33d0d929e755196e92c8e7c26f24a38cc76dd17b0e78c304426e7e6d34e3f4e39f5dae1397c3cfac8238eb45533a62e118b183afd66b652fd1d03a@54.249.78.93:30311",
	"enode://c9d6f27e2a2a3d208e02e043106a7fdf94806fca59d50c292b9ae28958162bfc2090a0a38208a68dace3f21674238c27e5f3b34ed8614a7469c6322549305015@3.115.97.167:30311",
	"enode://e071da3e6d4ef916c25901e976da90cfc5ec7de3269de09719637787778824da3186638a0f94c03ad90cdafeab32263f6f6cd9b0fbed67a744c25d38f67a5b5e@54.65.247.12:30311",
	"enode://45146a7cb02cd127d21cf3c37f533623c5caf4dea31aecc619d5e47ccc38f8377f08be8ce60f9c3429efaa50825c435498880a46b4634ccdc3ed8eb72ba054ae@3.209.235.131:30311",
	"enode://590e3069563fa0c46f72b6724c028aa2dda8f4e9ca70b7f288b0c1dcf9d32d74ac98d05da5a3c6473f133b13ed98523c41a567e18e6eadc8b156a3a2b11a791c@3.212.19.242:30311",
	"enode://1cc4534b14cfe351ab740a1418ab944a234ca2f702915eadb7e558a02010cb7c5a8c295a3b56bcefa7701c07752acd5539cb13df2aab8ae2d98934d712611443@100.24.179.110:30311",
	"enode://82efe88c5070a94e1cb7ec218f3e916a1621eac64067c6973e3103492aebc40d8febf810e12a786e316b8f2f02f35bf93ceaca69bb1852fbf8d5f345cd75f04e@35.75.44.49:30311",
	"enode://cb9d9f49c314188b2b9eb38572f16300c64ec2af30cfb044e27aad7f459e4ef52dc49ea2c31b78bebd8055b3bfaa94ddb6e96e75c0bd5928f5b7dc791d84371e@34.246.249.249:30311",
	"enode://1f14309a0e695df7afa03512e931d2782301f185d068a98ae9b24e6be88637ad7a7390b8ea074e6b8eb4b62e3f9826936cd7164b4ee4102b74f5b8fc357181fd@18.213.93.106:30311",

	// from: https://github.com/ledgerwatch/erigon/issues/6137#issuecomment-1330266787
	"enode://f3cfd69f2808ef64838abd8786342c0b22fdd28268703c8d6812e26e109f9a7cb2b37bd49724ebb46c233289f22da82991c87345eb9a2dadeddb8f37eeb259ac@18.180.28.21:30311",
	"enode://c0e8d1abd27c3c13ca879e16f34c12ffee936a7e5d7b7fb6f1af5cc75c6fad704e5667c7bbf7826fcb200d22b9bf86395271b0f76c21e63ad9a388ed548d4c90@54.65.247.12:30311",
	"enode://b5772a14fdaeebf4c1924e73c923bdf11c35240a6da7b9e5ec0e6cbb95e78327690b90e8ab0ea5270debc8834454b98eca34cc2a19817f5972498648a6959a3a@54.170.158.102:30311",
	"enode://352a361a9240d4d23bb6fab19cc6dc5a5fc6921abf19de65afe13f1802780aecd67c8c09d8c89043ff86947f171d98ab06906ef616d58e718067e02abea0dda9@79.125.105.65:30311",
	"enode://21530e423b42aed17d7eef67882ebb23357db4f8b10c94d4c71191f52955d97dc13eec03cfeff0fe3a1c89c955e81a6970c09689d21ecbec2142b26b7e759c45@54.216.119.18:30311",
	"enode://d61a31410c365e7fcd50e24d56a77d2d9741d4a57b295cc5070189ad90d0ec749d113b4b0432c6d795eb36597efce88d12ca45e645ec51b3a2144e1c1c41b66a@34.204.129.242:30311",
	"enode://4653bc7c235c3480968e5e81d91123bc67626f35c207ae4acab89347db675a627784c5982431300c02f547a7d33558718f7795e848d547a327abb111eac73636@54.144.170.236:30311",
	"enode://1479af5ea7bda822e8747d0b967309bced22cad5083b93bc6f4e1d7da7be067cd8495dc4c5a71579f2da8d9068f0c43ad6933d2b335a545b4ae49a846122b261@52.7.247.132:30311",
}

var ChapelBootnodes = []string{}
var ChapelStaticPeers = []string{
	"enode://69a90b35164ef862185d9f4d2c5eff79b92acd1360574c0edf36044055dc766d87285a820233ae5700e11c9ba06ce1cf23c1c68a4556121109776ce2a3990bba@52.199.214.252:30311",
	"enode://330d768f6de90e7825f0ea6fe59611ce9d50712e73547306846a9304663f9912bf1611037f7f90f21606242ded7fb476c7285cb7cd792836b8c0c5ef0365855c@18.181.52.189:30311",
	"enode://df1e8eb59e42cad3c4551b2a53e31a7e55a2fdde1287babd1e94b0836550b489ba16c40932e4dacb16cba346bd442c432265a299c4aca63ee7bb0f832b9f45eb@52.51.80.128:30311",
	"enode://0bd566a7fd136ecd19414a601bfdc530d5de161e3014033951dd603e72b1a8959eb5b70b06c87a5a75cbf45e4055c387d2a842bd6b1bd8b5041b3a61bab615cf@34.242.33.165:30311",
	"enode://604ed87d813c2b884ff1dc3095afeab18331f3cc361e8fb604159e844905dfa6e4c627306231d48f46a2edceffcd069264a89d99cdbf861a04e8d3d8d7282e8a@3.209.122.123:30311",
	"enode://4d358eca87c44230a49ceaca123c89c7e77232aeae745c3a4917e607b909c4a14034b3a742960a378c3f250e0e67391276e80c7beee7770071e13f33a5b0606a@52.72.123.113:30311",
}

var RialtoBootnodes = []string{}
var RialtoStaticPeers = ChapelStaticPeers

var SokolBootnodes = []string{
	"enode://f11a0f80939b49a28bf99581da9b351a592ec1504b9d32a7dfda79b36510a891e96631239c4166e5c73368c21e9bb3241e7fd6929b899772e5a8fe9a7b7c3af6@45.77.52.149:30303",
	"enode://e08adce358fc26dfbe1f24ee578dceaa29575ca44a39d9041203131db5135aceba6241840a9b57b1540eeaf7b4eff1aead28a74641be43342c35af454abb31b3@199.247.18.10:30313",
	"enode://f1a5100a81cb73163ae450c584d06b1f644aa4fad4486c6aeb4c384b343c54bb66c744aa5f133af66ea1b25f0f4a454f04878f3e96ee4cd2390c047396d6357b@209.97.158.4:30303",
	"enode://f11a0f80939b49a28bf99581da9b351a592ec1504b9d32a7dfda79b36510a891e96631239c4166e5c73368c21e9bb3241e7fd6929b899772e5a8fe9a7b7c3af6@45.77.52.149:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

var MumbaiBootnodes = []string{
	"enode://320553cda00dfc003f499a3ce9598029f364fbb3ed1222fdc20a94d97dcc4d8ba0cd0bfa996579dcc6d17a534741fb0a5da303a90579431259150de66b597251@54.147.31.250:30303",
	"enode://f0f48a8781629f95ff02606081e6e43e4aebd503f3d07fc931fad7dd5ca1ba52bd849a6f6c3be0e375cf13c9ae04d859c4a9ae3546dc8ed4f10aa5dbb47d4998@34.226.134.117:30303",
}

var BorMainnetBootnodes = []string{
	"enode://0cb82b395094ee4a2915e9714894627de9ed8498fb881cec6db7c65e8b9a5bd7f2f25cc84e71e89d0947e51c76e85d0847de848c7782b13c0255247a6758178c@44.232.55.71:30303",
	"enode://88116f4295f5a31538ae409e4d44ad40d22e44ee9342869e7d68bdec55b0f83c1530355ce8b41fbec0928a7d75a5745d528450d30aec92066ab6ba1ee351d710@159.203.9.164:30303",
}

var GnosisBootnodes = []string{
	"enode://6765fff89db92aa8d923e28c438af626c8ae95a43093cdccbd6f550a7b6ce6ab5d1a3dc60dd79af3e6d2c2e6731bae629f0e54446a0d9da408c4eca7ebcd8485@3.75.159.31:30303",
	"enode://9a7c98e8ee8cdd3199db68092b48868847d4743a471b26afc2ff878bafaa829ed43ee405f9aff58ae13fce53b898f7c2e3c30cb80af8eb111682c3c13f686dbb@18.198.130.54:30303",
	"enode://2c4307831914c237801993eac4f9596d8b2f78e1e76830419b64cb23f0933e52cb1e2bb3009cb4af76454bb5bc296135b36869fd6c13e2c2e536a0780e60fe82@3.64.242.196:30303",
	"enode://074f68e1a7df5b0859314ff721d55b59d9690e93249c941660609a29b302f02864df4f93ee48884f7ede57dc7f7646379d017a43c9745e34baff049749896b50@3.126.169.151:30303",
	"enode://d239697375d7586c7ea1de790401c310b0b1d389326849fa3b7c7005833c7a6b9020e49dfb3b61abfa39135237ffc4ff219cb84ca7653069e8548497527aa432@107.22.4.120:30303",
	"enode://d5852bf415d89b756faa809f4ff3f8beb661dc7d60cfb4a5542f9a5fcdf41e1ed0708a210db64b8c7ca32426e04ef0a50da58974124fdf562a8510314d11e28c@3.26.206.142:30303",
	"enode://01d372392bb22dd8a91f8b10b6bbb8d80d2dbe98d695801e0df9e4bd4825781df84bba88361f24d1b6580a61313f64e6cec82e8d842ad5f1b3d7cf8d6d132da7@15.152.45.82:30303",
	"enode://aee88e803b8e54925081957965b2527961cd90f4d6d14664884580b429da44729678a1258a8b49a42d1582c9c7c5ded05733622f7ab442ad9c6f655545a5ecdd@54.207.220.169:30303",
}

var ChiadoBootnodes = []string{
	"enode://7dd44af6138120f328bb031eb56e00985c149319d4f1e33275b30be7fddadd8ccd9f7b9c3b35a16136a61e85b2b2d1de073f30ec1d0ddf576a33be8ff48d88d0@139.144.26.89:30303",
	"enode://317b9cee65ccf1d747b00e604242bfa3ae367beee8f149e28c5b2b88820f855ea7b5a75eb5327cfc3d8ca97adbf71538468290a46592ed7009f3fb394ec752f1@139.144.26.115:30303",
	"enode://b77ae97906155ebbb83fd32c87ab0aa57372a24abbd8aa4bae679f048b726de4a195709f613be4981e44b24640bc89e4824427d94e9a37afc148da8250c8ab2d@139.144.26.101:30303",
	"enode://69f8abfa3b0221161f8c19014b90857a18742554af27af73fd779c486728750a0ff11b873975f104fc5276a3a7c3b5b68cb3c26c815e9f78462901895d652124@139.144.26.85:30303",
	"enode://ac7fc76f9b2ab343fb2d091365a7f46d17018e525cbedfbf24b247c76657e934ef4df61cc2f6dad6bfcf722425e03e1a8a6e4e4b52743acc2319cb8ebf27d742@170.187.154.239:30303",
	"enode://712144ac396fd2298b3e2559e2930d7f3a36fded3addd66955224958f1845634067717ab9522757ed2948f480fc52add5676487c8378e9011a7e2c0ac2f36cc3@3.71.132.231:30303",
	"enode://595160631241ea41b187b85716f9f9572a266daa940d74edbe3b83477264ce284d69208e61cf50e91641b1b4f9a03fa8e60eb73d435a84cf4616b1c969bc2512@3.69.35.13:30303",
}

// PulsechainBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Pulsechain network.
var PulsechainBootnodes = []string{
	// Pulsechain Go Bootnodes
}

// PulsechainTestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Pulsechain testnet network.
var PulsechainTestnetBootnodes = []string{
	"enode://5942169e5173992b2bab93e36bb2773e82b0fe91f2e70239a48232e85da6023673e2fab608f5b7fad4b8dcd2a29c4f2a6f800522aa00262a1b371fb80c7ec620@3.236.202.85:30303",   // bootnode-aws-us-east-1-001
	"enode://fa1420f97362e6c3e86dfcd38c3877748860b51c021de2662e46daf007d6de22c7668a838c9bf3468496fd10c2374544d16911dd2e2554aef7981e2858349952@44.202.85.131:30303",  // bootnode-aws-us-east-1-002
	"enode://eee9a1665c202b7fec55be2a5c1106b283f5694a426d727215d3cb0287074d3c72bc7943282b2492b0d39246daad28032d09a2adbeddd43a3e7d4da0e40fd840@174.129.96.223:30303", // bootnode-aws-us-east-1-003
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis libcommon.Hash, networkID uint64, protocol string) string {
	var net string
	var dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"
	var tld = ".ethdisco.net"

	if networkID == NetworkIDByChainName(networkname.PulsechainChainName) || networkID == NetworkIDByChainName(networkname.PulsechainTestnetChainName) {
		tld = ".pulsedisco.net"
		dnsPrefix = "enrtree://APFXO36RU3TWV7XFGWI2TYF5IDA3WM2GPTRL3TCZINWHZX4R6TAOK@"
	}

	switch genesis {
	case MainnetGenesisHash:
		switch networkID {
		case NetworkIDByChainName(networkname.PulsechainTestnetChainName):
			net = "testnet-v3"
		default:
			net = "mainnet"
		}
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + tld
}

func BootnodeURLsOfChain(chain string) []string {
	switch chain {
	case networkname.MainnetChainName:
		return MainnetBootnodes
	case networkname.SepoliaChainName:
		return SepoliaBootnodes
	case networkname.RinkebyChainName:
		return RinkebyBootnodes
	case networkname.GoerliChainName:
		return GoerliBootnodes
	case networkname.BSCChainName:
		return BscBootnodes
	case networkname.ChapelChainName:
		return ChapelBootnodes
	case networkname.RialtoChainName:
		return RialtoBootnodes
	case networkname.SokolChainName:
		return SokolBootnodes
	case networkname.MumbaiChainName:
		return MumbaiBootnodes
	case networkname.BorMainnetChainName:
		return BorMainnetBootnodes
	case networkname.GnosisChainName:
		return GnosisBootnodes
	case networkname.ChiadoChainName:
		return ChiadoBootnodes
	case networkname.PulsechainChainName:
		return PulsechainBootnodes
	case networkname.PulsechainTestnetChainName:
		return PulsechainTestnetBootnodes
	default:
		return []string{}
	}
}

func StaticPeerURLsOfChain(chain string) []string {
	switch chain {
	case networkname.BSCChainName:
		return BscStaticPeers
	case networkname.ChapelChainName:
		return ChapelStaticPeers
	case networkname.RialtoChainName:
		return RialtoStaticPeers
	case networkname.SepoliaChainName:
		return SepoliaStaticPeers
	default:
		return []string{}
	}
}
