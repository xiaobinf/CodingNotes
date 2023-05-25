# 测试开发Guide|测试岗位面试题汇总

# 前言(善用Ctrl+F)

![](http://tva1.sinaimg.cn/large/006F2AR3gy1gdpsdb0u7nj31062nxti1.jpg)



# 目录

### 一、开场白

> Q：简单自我介绍一下吧
>
> Q：项目和实习经历？（注意使用STAR法则表述）
>

#### star法则

STAR 法则是一种用于回答面试问题的技巧，它可以帮助您清晰、有条理地回答问题，向面试官展示您的能力和经验。STAR 是 Situation（情境）、Task（任务）、Action（行动）和Result（结果）的缩写。

以下是每个部分的详细说明：

Situation（情境）：首先，您需要描述一个具体的情境或场景，例如您所面临的挑战、问题或机会。这可以帮助面试官了解您所面对的情况，并为您的回答提供背景信息。

Task（任务）：接下来，您需要描述您所面临的具体任务或目标。这可以帮助面试官了解您所需要完成的工作，并为您的回答提供重点。

Action（行动）：然后，您需要描述您所采取的具体行动或步骤，以完成任务或达成目标。这可以帮助面试官了解您的能力和经验，并为您的回答提供具体的细节。

Result（结果）：最后，您需要描述您所取得的具体结果或成果。这可以帮助面试官了解您的成就和贡献，并为您的回答提供结论。

使用 STAR 法则回答问题时，您应该尽可能提供具体、量化的信息，例如具体的数字、时间表、成本节约等。这可以帮助面试官更好地了解您的能力和经验，并为您的回答提供更有说服力的证据。

###  二、软件测试基础

#### 工作内容

> Q：按测试内容划分，测试有哪些种类？

按测试内容划分，测试可以分为以下几种类型：

功能测试：测试软件的功能是否符合需求规格说明书中的要求，包括输入、输出、处理和操作等方面。

性能测试：测试软件在不同负载下的性能表现，包括响应时间、吞吐量、并发性等方面。

安全测试：测试软件的安全性，包括漏洞、攻击、数据泄露等方面。

兼容性测试：测试软件在不同操作系统、浏览器、设备等环境下的兼容性，包括界面、功能、性能等方面。

可靠性测试：测试软件的可靠性和稳定性，包括错误处理、恢复能力、容错性等方面。

可用性测试：测试软件的易用性和用户体验，包括界面设计、操作流程、反应速度等方面。

安装测试：测试软件的安装和卸载过程，包括安装、配置、升级、卸载等方面。

国际化测试：测试软件在不同语言、文化、地区等环境下的适应性，包括界面、文本、日期、货币等方面。

接口测试：测试软件与其他系统或组件之间的接口是否正常，包括数据传输、格式、协议等方面。

非功能性测试：测试软件的非功能性要求，包括可维护性、可扩展性、可重用性、可测试性等方面。

以上是按测试内容划分的一些常见的测试类型，不同类型的测试可以针对不同的测试目标和测试需求进行选择和组合。
 
> Q：软件开发（测试）的流程是怎么样的？

软件开发（测试）的流程通常包括以下几个阶段：

需求分析阶段：在这个阶段，开发团队与客户或用户沟通，收集和分析需求，确定软件的功能和性能要求，编写需求规格说明书。

设计阶段：在这个阶段，开发团队根据需求规格说明书，设计软件的架构、模块、接口和数据结构等，编写设计文档。

编码阶段：在这个阶段，开发团队根据设计文档，编写软件代码，进行单元测试和集成测试，确保代码的正确性和质量。

测试阶段：在这个阶段，测试团队根据测试计划和测试用例，对软件进行各种类型的测试，包括功能测试、性能测试、安全测试、兼容性测试等，发现和修复缺陷。

部署阶段：在这个阶段，开发团队将软件部署到生产环境中，进行系统测试和用户验收测试，确保软件的稳定性和可用性。

维护阶段：在这个阶段，开发团队对软件进行维护和升级，包括修复缺陷、添加新功能、优化性能等，确保软件的持续运行和发展。

以上是软件开发（测试）的一般流程，不同的项目和组织可能会有所不同，但是这些阶段通常都是必不可少的。在每个阶段中，开发团队和测试团队需要密切合作，及时沟通和协调，确保软件的质量和进度。

> Q：你印象最深刻的 bug 是？
>
> Q：谈谈你对 CI/CD 的理解

CI/CD 是一种软件开发和交付的方法论，它的目标是通过自动化和持续集成、持续交付、持续部署等技术手段，实现软件开发和交付的高效、快速和可靠。

具体来说，CI/CD 包括以下几个方面：

持续集成（Continuous Integration，CI）：将开发人员的代码集成到共享代码库中，并自动构建和测试代码，以确保代码的质量和稳定性。

持续交付（Continuous Delivery，CD）：将经过测试和验证的代码自动部署到预生产环境中，以便进行进一步的测试和验收。

持续部署（Continuous Deployment，CD）：将经过测试和验证的代码自动部署到生产环境中，以便实现快速、可靠的软件交付。

CI/CD 的优点包括：

提高软件开发和交付的效率和质量，减少手动操作和人为错误。

加快软件交付的速度和频率，缩短上线时间和反馈周期。

提高软件的可靠性和稳定性，减少故障和风险。

改善开发团队和运维团队之间的协作和沟通，促进团队的合作和创新。

总之，CI/CD 是一种现代化的软件开发和交付方法，它可以帮助开发团队和运维团队更好地协作和交付软件，提高软件的质量和效率，满足用户的需求和期望。

> Q：谈谈你对 DevOps 的理解

DevOps 是一种软件开发和运维的方法论，它的目标是通过自动化和协作，实现软件开发和运维的高效、快速和可靠。

具体来说，DevOps 包括以下几个方面：

自动化：通过自动化工具和流程，实现软件开发、测试、部署和运维的自动化，减少手动操作和人为错误。

协作：通过团队协作和沟通，实现开发团队和运维团队之间的协作和共享，促进团队的合作和创新。

持续交付：通过持续集成、持续交付和持续部署等技术手段，实现软件的快速、可靠的交付和部署。

监控和反馈：通过监控和反馈机制，实时监测软件的运行状态和性能指标，及时发现和解决问题，提高软件的可靠性和稳定性。

DevOps 的优点包括：

提高软件开发和运维的效率和质量，减少手动操作和人为错误。

加快软件交付的速度和频率，缩短上线时间和反馈周期。

提高软件的可靠性和稳定性，减少故障和风险。

改善开发团队和运维团队之间的协作和沟通，促进团队的合作和创新。

总之，DevOps 是一种现代化的软件开发和运维方法，它可以帮助开发团队和运维团队更好地协作和交付软件，提高软件的质量和效率，满足用户的需求和期望。

> Q: CI/CD和DEVOPS的关系
CICD（Continuous Integration/Continuous Delivery）和DevOps是紧密相关的概念，它们之间有着密切的关系。

CICD是一种软件开发和交付的方法论，它强调通过自动化和持续集成、持续交付、持续部署等技术手段，实现软件开发和交付的高效、快速和可靠。CICD的目标是将软件开发和交付的过程自动化，减少手动操作和人为错误，提高软件的质量和效率。

DevOps是一种软件开发和运维的方法论，它强调通过自动化和协作，实现软件开发和运维的高效、快速和可靠。DevOps的目标是将软件开发和运维的过程整合起来，促进开发团队和运维团队之间的协作和共享，提高软件的可靠性和稳定性。

CICD和DevOps的关系在于，CICD是实现DevOps的一种技术手段。CICD强调自动化和持续集成、持续交付、持续部署等技术手段，可以帮助开发团队和运维团队更好地协作和交付软件，实现DevOps的目标。同时，DevOps强调协作和共享，可以促进CICD的实现和应用，提高软件的质量和效率。

因此，CICD和DevOps是相互依存、相互促进的概念，它们的目标都是实现软件开发和交付的高效、快速和可靠。

> Q：什么是 BDD ? 什么是 TDD ?

BDD（Behavior-Driven Development）和TDD（Test-Driven Development）都是软件开发中常用的测试方法。

BDD是一种基于行为的测试方法，它强调测试用例应该从用户的行为和需求出发，以用户故事和场景为基础，编写测试用例和测试代码。BDD的目标是帮助开发团队和业务团队更好地沟通和理解需求，以及更好地测试和验证软件的功能和性能。

TDD是一种基于测试的开发方法，它强调在编写代码之前，先编写测试用例，然后编写代码来满足测试用例的要求。TDD的目标是帮助开发团队更好地理解需求和设计，以及更好地测试和验证代码的正确性和质量。

两者的区别在于，BDD更加强调测试用例的编写应该从用户的行为和需求出发，以用户故事和场景为基础，更加注重测试用例的可读性和可维护性；而TDD更加强调测试用例的编写应该在编写代码之前，以测试驱动开发的方式来设计和实现代码，更加注重测试用例的覆盖率和正确性。

总之，BDD和TDD都是软件开发中常用的测试方法，它们都强调测试在软件开发中的重要性，可以帮助开发团队更好地理解需求和设计，以及更好地测试和验证软件的功能和性能。

> Q：APP的一个页面，你怎么区分是原生Native页面，还是H5？
- 看断网的情况。正常打开就是原生的。 404就是H5
- 安卓手机可以打开开发者选项中的显示布局边界，页面元素很多的情况下布局是一整块的是h5的，布局密密麻麻的是原生控件。页面有布局的是原生的，否则为h5页面。（仅针对安卓手机适用）。
- 复制文章的提示。长按页面试试，如果出现文字选择、粘贴功能的是H5页面，否则是native原生的页面。
- 页面加载方式。新页面导航栏下面有一条加载的线的话，这个页面就是H5页面，如果没有就是原生的。或者出现白屏的话，大概率是的。
- app顶部 导航栏是否会有关闭的操作。如果APP顶部导航栏当中出现了关闭按钮或者有关闭的图标，那么当前的页面肯定的H5，原生的不会出现（除非设计开发者故意弄的）
- 判断页面 下拉刷新。如果界面没有明显刷新现象的是原生的，如果有明显刷新现象（比如闪一下）的是H5页面（ios和android）。
- 下拉页面的时候显示网址提供方的一定是H5。
#### 测试方法

> Q：黑盒测试的方法有哪些？

黑盒测试是一种软件测试方法，它不考虑软件的内部结构和实现细节，只关注软件的输入和输出，以及软件是否符合规格说明书和用户需求。常用的黑盒测试方法包括以下几种：

等价类划分法：将输入数据划分为若干个等价类，每个等价类代表一组具有相同特征的输入数据，然后从每个等价类中选择一个或多个测试用例进行测试。

边界值分析法：将输入数据的边界值作为测试用例，测试软件在边界值处的行为和响应。

决策表测试法：将软件的决策表转换为测试用例，测试软件在不同条件下的行为和响应。

因果图测试法：将软件的因果图转换为测试用例，测试软件在不同条件下的行为和响应。

状态转换测试法：将软件的状态转换图转换为测试用例，测试软件在不同状态下的行为和响应。

错误推测测试法：根据常见的错误和异常情况，推测软件可能出现的错误和异常情况，然后编写测试用例进行测试。

以上是常用的黑盒测试方法，不同的测试方法适用于不同的测试场景和测试目的，测试人员可以根据具体情况选择合适的测试方法进行测试。

> Q：白盒测试的方法有哪些？

白盒测试是一种软件测试方法，它考虑软件的内部结构和实现细节，以代码为基础，通过检查代码的逻辑和执行路径，来测试软件的正确性和质量。常用的白盒测试方法包括以下几种：

语句覆盖测试：测试用例覆盖代码中的每个语句至少一次。

判定覆盖测试：测试用例覆盖代码中的每个判定条件的真假至少一次。

条件覆盖测试：测试用例覆盖代码中的每个条件的真假至少一次。

路径覆盖测试：测试用例覆盖代码中的每个可能的执行路径至少一次。

边界值测试：测试用例覆盖代码中的边界值情况，例如最大值、最小值、空值等。

异常处理测试：测试用例覆盖代码中的异常处理情况，例如输入错误、文件不存在等。

性能测试：测试软件的性能和资源消耗情况，例如响应时间、并发用户数等。

以上是常用的白盒测试方法，不同的测试方法适用于不同的测试场景和测试目的，测试人员可以根据具体情况选择合适的测试方法进行测试。

> Q：什么是单元测试？

单元测试是一种软件测试方法，它是针对软件中最小的可测试单元——函数、方法或类等进行测试的过程。单元测试的目的是验证每个单元的功能是否正确，以及是否符合预期的行为和输出。

单元测试通常由开发人员编写，它可以在开发过程中及时发现和解决代码中的问题，提高代码的质量和可维护性。单元测试可以帮助开发人员更好地理解需求和设计，以及更好地测试和验证代码的正确性和质量。

单元测试的特点包括：

自动化：单元测试通常使用自动化测试工具和框架进行测试，可以快速、准确地执行测试用例，减少手动操作和人为错误。

独立性：单元测试通常是针对单个函数、方法或类进行测试，与其他模块和组件相互独立，可以更好地定位和解决问题。

可重复性：单元测试的测试用例可以重复执行，可以验证代码的正确性和稳定性。

可维护性：单元测试可以帮助开发人员更好地理解代码的结构和实现，以及更好地维护和更新代码。

总之，单元测试是一种重要的软件测试方法，它可以帮助开发人员更好地测试和验证代码的正确性和质量，提高代码的质量和可维护性。

> Q：什么是集成测试？

集成测试是一种软件测试方法，它是在单元测试之后，将多个单元组合在一起进行测试的过程。集成测试的目的是验证不同单元之间的交互和协作是否正确，以及是否符合预期的行为和输出。

集成测试通常由测试人员或测试团队进行，它可以在软件开发的不同阶段进行，例如模块集成测试、系统集成测试、验收测试等。集成测试可以帮助测试人员发现和解决不同单元之间的问题，以及验证软件的整体功能和性能。

集成测试的特点包括：

逐步集成：集成测试通常采用逐步集成的方式，从少量单元开始，逐步增加单元的数量和复杂度，直到整个系统被完全集成。

接口测试：集成测试通常关注不同单元之间的接口和交互，测试单元之间的数据传递、函数调用、消息传递等。

功能测试：集成测试也包括对整个系统的功能进行测试，验证系统是否符合预期的行为和输出。

性能测试：集成测试也可以包括对系统的性能和负载进行测试，例如响应时间、并发用户数等。

总之，集成测试是一种重要的软件测试方法，它可以帮助测试人员发现和解决不同单元之间的问题，以及验证软件的整体功能和性能。集成测试可以在软件开发的不同阶段进行，以确保软件的质量和稳定性。

> Q：测试用例怎么编写与设计？

测试用例是软件测试的基本单元，它描述了一个或多个测试场景，以及预期的行为和输出。测试用例的编写和设计是软件测试的重要环节，下面是一些编写和设计测试用例的基本步骤：

确定测试目标：测试用例的编写和设计应该始于测试目标的明确和清晰。测试目标应该包括测试的范围、测试的目的、测试的重点等。

确定测试场景：测试用例应该基于测试场景进行设计，测试场景应该包括输入、操作、预期输出等。

设计测试用例：根据测试场景，设计测试用例，测试用例应该包括测试步骤、输入数据、预期输出等。

确定测试数据：测试用例的设计需要确定测试数据，测试数据应该包括正常数据、边界数据、异常数据等。

确定测试环境：测试用例的设计需要确定测试环境，测试环境应该包括硬件、软件、网络等。

确定测试方法：测试用例的设计需要确定测试方法，测试方法应该包括手动测试、自动化测试等。

编写测试用例：根据测试场景和测试目标，编写测试用例，测试用例应该清晰、简洁、易于理解和执行。

执行测试用例：执行测试用例，记录测试结果，包括测试通过、测试失败、测试错误等。

分析测试结果：分析测试结果，确定问题的原因和解决方案，以及是否需要重新设计和编写测试用例。

总之，测试用例的编写和设计是软件测试的重要环节，它需要根据测试目标、测试场景、测试数据、测试环境、测试方法等因素进行设计和编写。测试用例应该清晰、简洁、易于理解和执行，以确保测试的有效性和可靠性。

> Q：什么是灰盒测试？

灰盒测试是介于黑盒测试和白盒测试之间的一种测试方法，它既考虑了系统的功能和行为，又考虑了系统的内部结构和实现。灰盒测试通常是在有限的代码访问权限下进行的，例如可以查看代码、调试代码、修改配置文件等。

灰盒测试的特点包括：

既考虑功能又考虑结构：灰盒测试既考虑系统的功能和行为，又考虑系统的内部结构和实现，可以更全面地测试系统的质量和稳定性。

有限的代码访问权限：灰盒测试通常是在有限的代码访问权限下进行的，例如可以查看代码、调试代码、修改配置文件等，可以更好地理解和测试系统的内部结构和实现。

可重复性：灰盒测试的测试用例可以重复执行，可以验证代码的正确性和稳定性。

可维护性：灰盒测试可以帮助测试人员更好地理解代码的结构和实现，以及更好地维护和更新代码。

灰盒测试通常包括以下测试方法：

分支覆盖测试：测试人员通过查看代码，测试每个分支的执行情况，以确保代码的完整性和正确性。

代码审查：测试人员通过查看代码，发现和解决代码中的问题，以确保代码的质量和可维护性。

数据库测试：测试人员通过查看数据库结构和数据，测试系统的数据完整性和正确性。

总之，灰盒测试是一种介于黑盒测试和白盒测试之间的测试方法，它既考虑了系统的功能和行为，又考虑了系统的内部结构和实现。灰盒测试可以更全面地测试系统的质量和稳定性，以确保系统的正确性和可靠性。

#### 测试文档
<details>
<summary>测试用例都包含哪些要素？</summary>
测试用例是用来验证软件系统是否符合需求和规格说明的一组测试步骤和输入数据。测试用例通常包含以下要素：

测试用例编号：用于标识测试用例的唯一编号。

测试用例名称：用于描述测试用例的名称，通常包含被测试的功能或模块名称。

测试用例描述：用于描述测试用例的目的、输入数据、预期结果等信息。

测试步骤：用于描述测试用例的具体步骤，包括输入数据、操作步骤、预期结果等信息。

输入数据：用于描述测试用例的输入数据，包括输入数据的类型、格式、范围等信息。

预期结果：用于描述测试用例的预期结果，包括输出数据的类型、格式、范围等信息。

实际结果：用于记录测试用例的实际结果，包括输出数据的类型、格式、范围等信息。

测试结果：用于记录测试用例的测试结果，包括测试通过、测试失败、测试未完成等信息。

测试人员：用于记录测试用例的测试人员，包括测试人员的姓名、工号等信息。

测试日期：用于记录测试用例的测试日期，包括测试开始时间、测试结束时间等信息。

总之，测试用例是用来验证软件系统是否符合需求和规格说明的一组测试步骤和输入数据，包含测试用例编号、测试用例名称、测试用例描述、测试步骤、输入数据、预期结果、实际结果、测试结果、测试人员和测试日期等要素。
</details>


<details>
<summary>测试报告需要展示哪些要素？</summary>
</details>
>
> Q：测试排期应该怎么估算？
>
> Q：谈谈你构造数据的经历？

#### 自动化测试

> Q：举例说明，都有哪些功能可以用自动化来进行？
>
> Q：如何判断一个功能能否进行自动化测试？
>
> Q：定位xpath路径都有哪些方法？

<details>
<summary>Q：如何定位一个动态的元素？</summary>
定位一个动态的元素，可以使用浏览器的开发者工具来进行查找和定位。具体步骤如下：

打开浏览器，进入需要定位元素的网页。

按下F12键或者右键选择“检查”打开浏览器的开发者工具。

在开发者工具中选择“元素”或“检查”选项卡。

使用鼠标在页面上点击需要定位的元素，开发者工具会自动定位到该元素的代码位置。

如果需要定位一个动态的元素，可以使用开发者工具中的“监视”或“网络”选项卡，查看元素的动态变化和网络请求。

如果元素是通过JavaScript代码动态生成的，可以使用开发者工具中的“调试”选项卡，设置断点、单步执行等方式进行调试。

如果元素是通过Ajax请求动态加载的，可以使用开发者工具中的“网络”选项卡，查看Ajax请求的响应内容和状态码。

总之，定位一个动态的元素，可以使用浏览器的开发者工具进行查找和定位，通过监视元素、查看网络请求、调试JavaScript代码等方式进行定位和调试。
</details>

<details>
<summary>Q：如何定位类似于悬浮在web页面上的元素（类似于web上飞来飞去的广告）</summary>
定位类似于悬浮在web页面上的元素，可以使用浏览器的开发者工具来进行查找和定位。具体步骤如下：

打开浏览器，进入需要定位元素的网页。

按下F12键或者右键选择“检查”打开浏览器的开发者工具。

在开发者工具中选择“元素”或“检查”选项卡。

使用鼠标在页面上点击需要定位的元素，开发者工具会自动定位到该元素的代码位置。

如果需要查找类似于悬浮在页面上的元素，可以使用开发者工具中的“选择元素”或“查找元素”功能，通过鼠标选择或输入元素的CSS选择器、XPath表达式等方式进行查找和定位。

如果需要调试类似于悬浮在页面上的元素的JavaScript代码，可以使用开发者工具中的“调试”选项卡，设置断点、单步执行等方式进行调试。

总之，定位类似于悬浮在web页面上的元素，可以使用浏览器的开发者工具进行查找和定位，通过选择元素、查找元素、调试JavaScript代码等方式进行定位和调试。
</details>

> Q：列举你知道的自动化测试工具

#### 测试工具
<details>
<summary>Q：聊聊 fiddler 的抓包原理</summary>
finddler是一个Web代理服务器，它可以捕获客户端和服务器之间的HTTP（S）请求。它的工作原理是：
finddler启动后，会自动将代理服务器设置为本机，端口为8888。
当客户端发起一个HTTP（S）请求时，finddler会拦截该请求，并伪装成客户端发送相同的请求给Web服务器。
Web服务器收到finddler的请求后，会返回响应给finddler。
finddler收到Web服务器的响应后，会保存并分析该响应，并伪装成Web服务器返回相同的响应给客户端。
这样，finddler就可以查看和修改客户端和服务器之间的所有数据包。
如果请求是HTTPS的，finddler还需要进行一些额外的步骤，如生成自签名证书，解密对称密钥等，以便能够解密加密的数据包。
</details>


<details>
<summary>Q：fiddler 怎么抓取 https 请求？</summary>
Fiddler是一款常用的网络抓包工具，它可以抓取HTTP和HTTPS请求并对其进行分析。要抓取HTTPS请求，需要进行以下设置：

安装Fiddler证书：Fiddler需要安装自己的根证书才能解密HTTPS流量。在Fiddler中选择“Tools”->“Options”->“HTTPS”选项卡，勾选“Decrypt HTTPS traffic”选项，并点击“Actions”->“Export Root Certificate to Desktop”按钮，将证书保存到桌面。

安装Fiddler证书到浏览器：将Fiddler证书导入到浏览器的受信任根证书中。在浏览器中打开证书文件，选择“安装证书”，选择“计算机账户”和“受信任的根证书颁发机构”，完成证书安装。

启用HTTPS代理：在Fiddler中选择“Tools”->“Options”->“Connections”选项卡，勾选“Allow remote computers to connect”和“Act as system proxy on startup”选项。

启用HTTPS解密：在Fiddler中选择“Tools”->“Options”->“HTTPS”选项卡，勾选“Decrypt HTTPS traffic”选项。

开始抓取HTTPS请求：在Fiddler中选择“File”->“Capture Traffic”选项，开始抓取HTTPS请求。

总之，要抓取HTTPS请求，需要安装Fiddler证书、安装证书到浏览器、启用HTTPS代理和启用HTTPS解密等步骤。这样就可以使用Fiddler抓取HTTPS请求并进行分析。
</details> 


<details>
<summary>Q：wireshark 抓包的头部有什么？</summary>
Wireshark是一款常用的网络抓包工具，它可以抓取网络数据包并对其进行分析。Wireshark抓包的头部包含以下信息：

Frame：帧头，包含帧的长度、时间戳、帧的序号等信息。

Ethernet：以太网头，包含源MAC地址、目的MAC地址、协议类型等信息。

IP：IP头，包含源IP地址、目的IP地址、协议类型、TTL等信息。

TCP/UDP：TCP或UDP头，包含源端口号、目的端口号、序号、确认号、标志位等信息。

HTTP/HTTPS：HTTP或HTTPS头，包含请求方法、请求URL、响应状态码、响应内容等信息。

DNS：DNS头，包含查询类型、查询名称、响应IP地址等信息。

SSL/TLS：SSL或TLS头，包含加密算法、证书信息等信息。

ICMP：ICMP头，包含类型、代码、校验和等信息。

总之，Wireshark抓包的头部包含了网络数据包的各种信息，包括帧头、以太网头、IP头、TCP/UDP头、HTTP/HTTPS头、DNS头、SSL/TLS头、ICMP头等。这些信息可以帮助用户了解网络数据包的来源、目的、协议类型、内容等信息，从而进行网络分析和故障排除。
</details>

#### 性能测试

> Q：性能测试如何做的？
>
> Q：性能测试需要关注哪些方面？

### 三、软件测试实战

#### 排查问题的思路

> Q：网页崩溃的原因是什么？
>
> Q：有个用户反馈上传头像失败，分析原因？
>
> Q：app闪退的原因？
>
> Q：偶然闪退的排查？
>
> Q：网页卡顿的原因是什么？
>
> Q：10%的用户反馈用不了功能，你讲如何排查？
>
> Q：登录的按钮不能点击，如何排查问题？
>
> Q：压测的时候，QPS一直上不去，你会怎么排查？
>
> Q：APP提示无法连接网络，你会如何排查？
>
> Q：怎么判断一个BUG到底是前端的BUG还是后端的BUG？

#### 实战案例

> Q：微博发动态，设计一下测试点
>
> Q：对一台自动售货机进行测试用例设计
>
> Q：设计微信发红包测试用例
>
> Q：设计抖音直播功能测试用例
>
> Q：设计微信扫码支付的测试用例
>
> Q：设计百度首页的测试用例
>
> Q：微信的点赞功能怎么测试？
>
> Q：微信红包是先计算每个人能获得的钱还是当这个人点了再计算。
>
> Q：微信朋友圈评论功能怎么测试？
>
> Q：微信上线一个新的好友推荐功能功能如何测试？
>
> Q：测试微信换头像功能，设计测试用例
>
> Q：抖音视频的安全性测试，测试点有哪些？
>
> Q：如果手机浏览器输入 http://baidu.com 打不开页面，你会怎么排查？
>
> Q：设计输入框测试用例？
>
> Q：编写一个登录界面的测试用例？
>
> Q：对一个接口编写测试用例
>
> Q：搜索功能怎么测试？

### 四、语言基础
#### Java

#### Python

> Q：全局变量和局部变量变量名能否一样？
>
> Q：Python 里 is 和 == 的区别？
>
> Q：Python 变量的创建与消亡过程
>
> Q：Python的垃圾回收的机制
>
> Q：dict的底层结构，tuple和list的底层结构的区别
>
> Q：深拷贝浅拷贝的区别是什么？
>
> Q：什么是协程？
>
> Q：什么是装饰器？举一个你用过装饰器的例子
>
> Q：@classmethod 和 @staticmethod 的区别，以及分别运用在哪些使用场景？
>
> Q：什么是鸭子类型？
>
> Q：python 的可变类型有哪些？
>
> Q：python 常见的数据类型有哪些？
>
> Q：列举 python2 和 python3 的区别
>
> Q：什么是lambda函数？怎么用？
>
> Q：别的编程语言都有三目运算符，三目运算符在python中怎么表达？
>
> Q：try...except...else 和 try...except...finally 的区别是什么？
>
> Q：什么是可迭代对象？可迭代对象的原理是什么？
>
> Q：with...as 的原理是什么？
>
> Q：解释一下python的GIL锁
>
> Q：python 是单继承还是多继承？
>
> Q：python 继承的顺序是什么？
>
> Q：什么是元类？
>
> Q：为什么都说 python 慢？

#### Shell

> Q：什么场景下，适合编写 Shell 脚本来处理？


#### C++
<details>
<summary>Q：指针和引用的区别是什么？</summary>
C++中的指针和引用都是用来处理内存地址的，但它们有以下几个区别：

定义方式不同：指针是一个变量，它存储了一个内存地址，可以通过“*”操作符来访问该地址上的值；引用是一个别名，它引用了一个已经存在的变量，可以直接访问该变量的值。

空指针和空引用不同：指针可以被赋值为空指针，即指向空地址；引用必须在定义时初始化，不能引用空值。

操作方式不同：指针可以被重新赋值，可以进行指针运算，可以指向不同类型的变量；引用一旦被初始化，就不能再引用其他变量，也不能进行引用运算。

传递方式不同：指针可以作为函数参数传递，可以通过指针修改函数外部的变量；引用也可以作为函数参数传递，可以通过引用修改函数外部的变量，但是引用更加简洁和安全。

总之，指针和引用都是用来处理内存地址的，但它们的定义方式、空值处理、操作方式和传递方式等方面有所不同。在使用时需要根据具体情况选择合适的方式，以确保程序的正确性和可靠性。
</details>

#### Golang
> Q: slice的底层实现

> Q: map的底层实现


### 五、数据结构与算法

#### 链表

> Q：一个有序链表，怎么求第K大个节点？
>
> Q：怎么找出这两个链表是否有相交的点
>
> Q：数组与链表的区别
>
> Q：链表逆序

#### 数组

> Q：怎么对俩有序数组合并？
>
> Q：求数组中和为 n 的两个数，时间复杂度是多少？
>
> Q：int 型数组，怎么排序？
>
> Q：int 型数组，怎么去重？

#### 复杂度

> Q：时间复杂度是什么？
>
> Q：怎么计算时间复杂度？

#### 排序

> Q：八大排序都有哪些？
>
> Q：快排的原理
>
> Q：归并排序的原理
>
> Q：冒泡排序的原理
>
> Q：插入排序的原理
>
> Q：选择排序的原理

#### 树

> Q：数的定义，代码实现
>
> Q：什么是树的高度？怎么求？
>
> Q：计算二叉树节点的个数
>
> Q：什么是根结点？什么是叶子节点？
>
> Q：打印二叉树

#### 递归

> Q：什么是递归？
>
> Q：递归的时间复杂度是什么？

#### 字符串

> Q：字符串长度可以改变么？



#### 堆与栈

> Q：堆和栈都有什么区别
> Q：代码实现栈


#### 动态规划

> Q：什么是动态规划？

#### 高级算法

> Q：什么是深度遍历？什么是广度遍历？

#### 查找

> Q：能够实现二分查找的必要条件是什么？

#### 哈希
> 待施工

#### 图
> 待施工

### 六、框架

#### 开发框架

##### Django

> Q：简单介绍一下Django的目录结构
>
> Q：了解什么是MVC么？为什么说 Django 是MTV模型？
>
> Q：一个网络请求在Django中的实现过程
>
> Q：uwsgi是什么？用uwsgi和用Django的原生启动方式，有什么区别？
>
> Q：了解 ORM 么？简单介绍一下Django的ORM映射
>
> Q：Django的 CBV 和 FBV 的区别？实际项目中你会使用哪种方式？为什么？
>
> Q：怎么把 Django 的 module 同步到数据库中？同步过程中会遇到什么坑么？

#### 测试框架

##### selenium

> Q：selenium框架的运行原理
>
> Q：selenium定位元素的方法都有哪些

##### Appium

> Q：Appium用过吗？原理是什么？

##### Unittest
> 待施工

##### Pytest
> 待施工

##### TestNG
> 待施工

##### Junit
> 待施工

### 七、计算机基础

#### 计算机网络

##### TCP/UDP

> Q：简单介绍 TCP 三次握手（为什么不是两次、四次）
>
> Q：四次分手是什么？
>
> Q：TCP 拥塞是什么？
>
> Q：TCP怎么保证安全的，UDP能否也像TCP那样安全，怎么做？
>
> Q：你知道 tcp 的控制可靠性的策略和重传机制么？
>
> Q：TCP协议属于哪一层？

##### HTTP/HTTPS

> Q：cookie 和 session 机制、区别
>
> Q：输入url到网页显示出来中间的过程
>
> Q：列举你知道的网页状态码
>
> Q：3 开头的网络状态码的含义是什么？302和304的区别是什么？
>
> Q：4 开头的网络状态码含义是什么？
>
> Q：5 开头的网络状态码含义是什么？
>
> Q：什么是 HTTPS？原理是什么？
>
> Q：GET 和 POST 的区别是什么？
>
> Q：网络请求 method 有哪几种？
>
> Q：简单介绍一下什么是 RESTful API
>
> Q：PUT 和 POST 的区别
>
> Q：列举常见的请求 Header 头
>
> Q：一个HTTP请求报文是什么样的？（GET举例）
>
> Q：GET 的长度限制了解么？

##### 其他

> Q：最大连接数和QPS区别
>
> Q：网络一共分为几层？
>
> Q：客户端向服务器请求图片和动态资源的区别
>
> Q：图片渲染的过程 前端
>
> Q：网络七层有哪些？tcp，udp，arp都在哪一层？
>
> Q：QPS 和 TPS 的区别是什么？
>
> Q：解释一下DNS
>
> Q：什么是反向代理？
>
> Q：什么是 socket？
>
> Q：【手撕】用 socket 实现一个聊天室功能？

#### 操作系统

##### 进程/线程

> Q：进程与线程的区别？
>
> Q：何时cpu处理进程最慢？
>
> Q：为什么会出现死锁？
>
> Q：进程间通信的方式？
>
> Q：线程间通信的方式？
>
> Q：如何做到线程同步？

##### Linux

> Q：linux切换目录
>
> Q：linux命令，统计一个文本中关键字出现的次数
>
> Q：linux 查找当前目录下所有后缀为 .py文件
>
> Q：知道的linux常用命令：查看指定端口进程
>
> Q：cd - 和 cd ～
>
> Q：linux 查看某个进程命令怎么写
>
> Q：如何查看日志？怎么查看后500条日志？
>
> Q：awk有什么用？如何用（举个例子）？
>
> Q：如何查看系统性能？
>
> Q：如何查看剩余磁盘空间大小？
>
> Q：如何查看目录占空间大小？
>
> Q：你知道 xarg 的用法吗？
>
> Q：怎么从本地计算机与服务器中进行文件传输？
>
> Q：测试服务器之间怎么进行文件拷贝？
>
> Q：sed 有什么用？如何用（举个例子）？
>
> Q：怎么杀死一个进程？
>
> Q：怎么递归删除一个目录下的所有文件？
>
> Q：怎么查看内存大小？
>
> Q：怎么查看 CPU 使用情况？
>
> Q：怎么重启 Linux 服务器？
>
> Q：怎么打印出一个文件的第500-1000行？

##### 文件存储

> Q：静态存储和动态存储的区别
>
> Q：视频在服务器的存储几种方式
>
> Q：CDN有什么用？
>
> Q：为什么前端静态资源要上传到CDN上？

##### 其他

> Q：系统资源包括哪些？
>
> Q：什么是 IO 操作？
>
> Q：什么是内存？
>
> Q：什么是硬盘？
>
> Q：什么是CPU？
>
> Q：什么操作比较消耗CPU？
>
> Q：什么是UTF-8？什么是Unicode？
>
> Q：什么是 IO 多路复用？以及怎么实现？
>
> Q：谈谈什么是分布式？为什么要用分布式？



#### 数据库

##### 非关系型数据库

> Q：非关系型数据库有哪些？
>
> Q：Redis的数据结构有哪些？
>
> Q：Redis 和 Memcached 的区别？
>
> Q：Redis 的用途？
>
> Q：Redis一秒能写入多少数据？
>
> Q：Redis为什么快？
>
> Q：Redis的过期时间怎么设置？哪些场景适合缓存更长时间？
>
> Q：什么是缓存雪崩？
>
> Q：什么是缓存击穿？缓存击穿和缓存穿透的区别是什么？
>
> Q：什么时候适合用MongoDB？
>
> Q：Redis 适合做消息队列吗？为什么？
>
> Q：Redis 的数据是存储在内存当中的，假如断电之后就会造成数据丢失，那怎么对 Redis 进行数据固化？

##### 关系型数据库

**数据库基本理论**

> Q：有哪些数据库优化的方式？
>
> Q：关系型数据库和非关系型数据库的区别？
>
> Q：数据库的事务有什么用？什么时候应该使用事务，什么时候不该使用事务？
>
> Q：数据库索引有什么用？什么是联合索引？
>
> Q：主键适合用自定义ID还是自动ID？

主键适合使用自动ID。

自动ID是指数据库系统自动生成的唯一标识符，通常使用自增长的整数或UUID等方式生成。自动ID的优点是简单、高效、可靠，可以确保每个记录都有唯一的标识符，避免了手动分配ID时可能出现的重复和错误。

相比之下，自定义ID需要开发人员手动分配，需要考虑到唯一性、长度、可读性等问题，容易出现错误和冲突。另外，自定义ID还需要考虑到分布式系统的情况，需要保证在不同的节点上生成的ID不会重复。

因此，一般情况下，主键适合使用自动ID。当然，在某些特殊情况下，例如需要将主键暴露给用户或需要保证主键的可读性时，可以考虑使用自定义ID。但是，需要注意到自定义ID可能会带来一些额外的开发和维护成本，需要根据实际情况进行选择。
![img.png](img.png)

> Q：外键是什么？为什么大公司有时候不建议使用外键？
>
> Q：索引设置得越多越好吗？索引得优缺点是什么？
>
> Q：InnoDB是基于什么实现的？
>
> Q：为什么会造成数据库死锁？怎么解决？
>
> Q：事务都有哪些特点？
>
> Q：数据库设计的三大范式是什么？
>
> Q：一张数据表最多不建议超过多少行？
>
> Q：为什么要进行分表分库？
>
> Q：分表分库有哪几种方式？举例说明

**SQL**

> Q：【手撕】数据库，查找一个学生两门功课都大于80分的姓名
>
> Q：【手撕】联表查询2个表中工号为“123”的人的所有信息
>
> Q：【手撕】一个人员表，一个部门表，人员表中存了部门id，查人员表各部门表所有数据
>
> Q：【手撕】查询一个城市列表里面重复的城市名，并且统计重复次数
>
> Q：【手撕】查找一个学生成绩表中平均分数大于90分的学生名单
>
> Q：【手撕】查找学生成绩表中平均成绩最高的同学
>
> Q：数据库怎么拷贝数据？
>
> Q：删除数据的方式有哪些？说说 drop table和truncate table的区别？


### 八、智力题

> Q：跳台阶问题
>
> Q：4分钟沙漏和7分钟沙漏怎么漏出9分钟
>
> Q：两个粗细不同的香，燃尽时间都是1个小时，怎么用这个2根香计算15分钟的时间
>
> Q：赛马
>
> Q：10堆苹果，每堆10个，9堆每个50g，1堆每个40g，有一个称，求只称一次，找出这个轻的一堆
>
> Q：飞机加油问题
>
> Q：逻辑：四个开关四个灯泡
>
> Q：地球弧形

### 九、编程题

> Q：求最大回文个数
>
> Q：一个数组中有正数有负数（没有0），请将它排成正负相间的数组（多余的全部放后面），时间复杂度
>
> 不超过O（n）;
>
> Q：一道编程题，输入一串由ABCD四个字随机组成的字符串和一个整数k，返回字符串种前k个字的顺序重复了几次。
>
> Q：编程：判断一个字符串是否符合ipv4格式
>
> Q：代码题，给一个句子，只把单词翻转然后输入
>
> Q：【手撕代码】字符串中只出现一次的字符、找出数组中最小的四个数字
>
> Q：代码：一串字符串中最小的整数
>
> Q：写代码，类似高考成绩，一个表中有很多数据（无序的），给你一个成绩，查出在表中的排名
>
> Q：编程题，鸡兔同笼，一半的兔子伸起一半的脚，输入地上有几只脚，列出所有的可能性（兔子是基数则整除2）
>
> Q：判断一个字符串是否是点分十进制的ipv4格式
>
> Q：100万个学生 按照成绩 及对应排名录入 分数查找排名（hashmap）
>
> Q：python的编程题，输入一个字符串然后空格切割在统计每个字母出现的次数

### 十、HR常问

> Q：为什么想做测试
>
> Q：对测开的理解
>
> Q：测试过程中有没有出现问题，是如何解决的
>
> Q：最近看了什么书？学了什么？为什么学？有看什么技术书籍吗？
>
> Q：个人优缺点，举例
>
> Q：测试看重什么能力
>
> Q：项目问题细挖
>
> Q：为什么选择xx公司？
>
> Q：你对我们公司有什么了解吗？
>
> Q：之前实习收获了什么
>
> Q：介绍下自己的优缺点
>
> Q：抗压能力如何，描述一件自己如何抗压的经历
>
> Q：反问环节：你有什么问题想问我么？
>
> Q：项目中收获了什么？
>
> Q：平时怎么学习的
>
> Q：为什么要离职？
>
> Q：你的期望薪资是多少？


# 关于我

 [【关于我】](http://mp.weixin.qq.com/s?__biz=MzAxMTA2OTY0Nw==&mid=301785273&idx=1&sn=feedd9980f4d5027e2bd2e4e59b2ef14&chksm=0f4ef1193839780f802d1ffa152a785682a514383ee27840727f17e9ecb8a783a558af50f6ac#rd)。